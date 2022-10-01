package main

import (
	"fmt"
	"main/core/routers"
)

import (
	"log"
	"main/core"
)

func main() {
	fmt.Println("Hello, World!")
	server, err := core.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	err, _ = routers.NewUser(server, "/v1/user")
	if err != nil {
		return
	}
}
