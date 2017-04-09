package main

import (
	"flag"
	"fmt"
	"time"
)

type server struct {
	port     string
	database map[string]string
}

func main() {
	fmt.Println("You are executing Server 2 at", time.Now())
	srv := &server{}
	srv.database = map[string]string{}
	flag.StringVar(&srv.port, "port", "50552", "port of grpc server")
	flag.Parse()
	srv.setupGrcpServer()
}
