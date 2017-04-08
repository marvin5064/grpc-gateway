package main

import (
	"fmt"
	"time"
)

type server struct{}

func main() {
	fmt.Println("You are executing Server at", time.Now())
	srv := &server{}
	srv.setupGrcpServer()
}
