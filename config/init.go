package config

import (
	"log"
	"math/rand"
	"os"
	"strconv"
)

var (
	serverPort uint16 = 8080
	brookPort  uint16
	brookPass  string
)

const (
	SERVER_PORT = "SERVER_PORT"
	BROOK_PORT  = "BROOK_PORT"
	BROOK_PASS  = "BROOK_PASS"
)

func genPasswd() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 6)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetBrookPass() string {
	return brookPass
}

func SetBrookPass(pass string) {
	brookPass = pass
}

func GetBrookPort() uint16 {
	return brookPort
}

func SetBrookPort(port uint16) {
	brookPort = port
}

func GetServerPort() uint16 {
	return serverPort
}

func Init() {
	// server port
	portStr := os.Getenv(SERVER_PORT)
	if len(portStr) > 0 {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			panic(err)
		}
		serverPort = uint16(port)
	}
	log.Printf("set server listening on port %v", serverPort)

	// brook port
	portStr = os.Getenv(BROOK_PORT)
	if len(portStr) == 0 {
		brookPort = uint16(rand.Int31n(20000) + 1025)
	} else {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			panic(err)
		}
		brookPort = uint16(port)
	}
	log.Printf("set brook listening on port %v", brookPort)

	// brook password
	brookPass = os.Getenv(BROOK_PASS)
	if len(brookPass) == 0 {
		brookPass = genPasswd()
		log.Printf("you don't set brook pass, generate default password %s", brookPass)
	}
}
