package main

import (
	"log"

	"github.com/marvin5064/grpc-gateway/protobuf/university"

	"google.golang.org/grpc"
)

func (c *client) setupUniversityMgrClient() {
	addr := "localhost:50551"
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln(err)
	}
	c.universityMgrCln = university.NewUniversityManagerClient(conn)
	log.Println("university manager client setup at", addr)
}
