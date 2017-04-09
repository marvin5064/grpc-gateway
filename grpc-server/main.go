package main

import (
	"flag"
	"fmt"
	"time"
)

type server struct {
	port string
}

func main() {
	fmt.Println("You are executing Server at", time.Now())
	srv := &server{}
	flag.StringVar(&srv.port, "port", "50551", "port of grpc server")
	flag.Parse()
	srv.setupGrcpServer()
}
