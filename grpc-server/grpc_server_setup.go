package main

import (
	"log"
	"net"

	"github.com/marvin5064/grpc-gateway/protobuf/university"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (s *server) setupGrcpServer() {
	addr := ":" + s.port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gsrv := grpc.NewServer()
	defer gsrv.Stop()
	university.RegisterUniversityManagerServer(gsrv, s)
	reflection.Register(gsrv)

	log.Println("setting up university manager grpc at", addr)
	err = gsrv.Serve(lis)
	if err != nil {
		log.Println(err)
	}
}
