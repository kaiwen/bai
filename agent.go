package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func spawn(port uint16, pass string) {
	// terminate old process
	file, err := os.OpenFile("/tmp/ab.pid", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	buf, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}
	if len(buf) != 0 {
		pid, err := strconv.Atoi(string(buf))
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("kill process %d", pid)
		process, _ := os.FindProcess(pid)
		process.Signal(os.Interrupt)
	}

	// start new process
	log.Printf("set port %d", port)
	process, err := os.StartProcess("/usr/local/bin/brook_linux_amd64", []string{}, &os.ProcAttr{})
	if err != nil {
		log.Println(err)
		return
	}
	process.Release()
}

func main() {
	log.Default().SetFlags(log.LstdFlags | log.Lshortfile)

	iport, err := strconv.Atoi(os.Args[1])
	port := uint16(iport)
	if err != nil {
		log.Fatal(err)
	}
	pass := os.Args[2]

	spawn(port, pass)

	http.HandleFunc("/port", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Write([]byte(strconv.Itoa(int(port))))
			return
		}
		if r.Method == "POST" {
			port_str := r.FormValue("port")
			iport, err := strconv.Atoi(port_str)
			if err != nil {
				log.Println(err)
				w.Write([]byte(err.Error()))
				return
			}
			port = uint16(iport)

			spawn(port, pass)

			w.Write([]byte("success"))
			return
		}
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
