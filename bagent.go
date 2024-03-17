package main

import (
	"log"
	"net/http"
	"strconv"

	"bagent.com/bagent/config"
	"bagent.com/bagent/controller"
	"bagent.com/bagent/service"
)

func main() {
	config.Init()
	controller.Init()

	if err := service.Respawn(config.GetBrookPort(), config.GetBrookPass()); err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println(http.ListenAndServe(":"+strconv.Itoa(int(config.GetServerPort())), nil))
}
