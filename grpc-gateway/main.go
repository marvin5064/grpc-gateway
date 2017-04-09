package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/marvin5064/grpc-gateway/protobuf/university"
	"google.golang.org/grpc"
)

func main() {
	mux := runtime.NewServeMux()
	err := university.RegisterUniversityManagerHandlerFromEndpoint(context.TODO(), mux, "localhost:50551", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("set up university manager gateway at :8080")
	http.ListenAndServe(":8080", mux)
}
