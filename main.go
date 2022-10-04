package main

import (
	"main/core/routers"
)

import (
	"log"
	"main/core"
)

func main() {
	server, err := core.NewServer()
	if err != nil {
		log.Fatal(err)
		return
	}

	err, _ = routers.NewUser(server, "")
	if err != nil {
		log.Fatal(err)
		return
	}

	err, _ = routers.NewMovie(server, "")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = server.Start()
	if err != nil {
		log.Fatal("start error", err)
		return
	}
}
