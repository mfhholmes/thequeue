package main

import (
	"fmt"
	"github.com/gocraft/web"
	"log"
	"net"
	"net/http"
)

type Server struct {
	Address string
	Port    int
	Static  string
	Lst     net.Listener
	Srv     http.Server
	Rooter  *web.Router
}

func (srv *Server) Setup(context interface{}) error {
	fmt.Println("server setting up")
	srv.Rooter = web.New(context)
	//set up static server
	fmt.Println("setting up static folder on " + srv.Static)
	srv.Rooter.Middleware(web.StaticMiddleware("./" + srv.Static))
	return nil
}

func (srv *Server) Start() error {
	port := fmt.Sprintf(":%d", srv.Port)
	fmt.Printf("server starting on %s \n", port)
	srv.Srv = http.Server{
		Addr:    port,
		Handler: srv.Rooter,
	}
	var err error
	srv.Lst, err = net.Listen("tcp", port)
	if nil != err {
		log.Printf("Listener failed to start because: %s", err.Error())
	}
	return err
}
