package service

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"

	"bagent.com/bagent/config"
	"bagent.com/bagent/utils"
)

var lock sync.Mutex

func Respawn(port uint16, pass string) error {
	// terminate old process
	file, err := os.OpenFile("/tmp/ab.pid", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return err
	}
	if len(buf) != 0 {
		pid, err := strconv.Atoi(string(buf))
		if err != nil {
			log.Println(err)
			return err
		}

		log.Printf("kill process %d", pid)
		process, _ := os.FindProcess(pid)
		process.Signal(os.Interrupt)
		process.Wait()
	}

	// start new process
	log.Printf("set port %d", port)
	cmd := exec.Command("/usr/local/bin/brook", "server", "-l", ":"+strconv.Itoa(int(port)), "-p", pass)
	if err := cmd.Start(); err != nil {
		log.Println(err)
		return err
	}
	file.Truncate(0)
	file.Seek(0, 0)
	file.WriteString(strconv.Itoa(cmd.Process.Pid))

	log.Println("detached process")
	return nil
}

func checkPass(pass string) bool {
	return pass == config.GetBrookPass()
}

func GetPort(pass string) (uint16, error) {
	if !checkPass(pass) {
		return 0, utils.ErrInvalidToken
	}
	return config.GetBrookPort(), nil
}

func SetPort(port uint16, pass string) error {
	lock.Lock()
	defer lock.Unlock()

	if !checkPass(pass) {
		return utils.ErrInvalidToken
	}
	if port == config.GetBrookPort() {
		return nil
	}

	config.SetBrookPort(port)
	return Respawn(port, pass)
}

func SetPass(newPass string, oldPass string) error {
	lock.Lock()
	defer lock.Unlock()

	if !checkPass(oldPass) {
		return utils.ErrInvalidToken
	}
	if oldPass == newPass {
		return nil
	}

	config.SetBrookPass(newPass)
	return Respawn(config.GetBrookPort(), newPass)
}

func SetPortAndPass(port uint16, newPass string, oldPass string) error {
	lock.Lock()
	defer lock.Unlock()

	if !checkPass(oldPass) {
		return utils.ErrInvalidToken
	}
	if oldPass == newPass && port == config.GetBrookPort() {
		return nil
	}

	config.SetBrookPort(port)
	config.SetBrookPass(newPass)
	return Respawn(port, newPass)
}
