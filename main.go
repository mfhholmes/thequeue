package main

import (
	"log"
)

var webserver = Server{
	Address: "localhost",
	Port:    8080,
}

func main() {
	ctx := Context{}
	err := webserver.Setup(ctx)
	if nil != err {
		log.Printf("Error setting up server: %s", err.Error())
		return
	}
	err = webserver.Start()
	if nil != err {
		log.Printf("Error starting server: %s", err.Error())
		return
	}
}

type Context struct {
}
