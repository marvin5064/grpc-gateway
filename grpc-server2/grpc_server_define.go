package main

import (
	"fmt"

	context "golang.org/x/net/context"

	"github.com/marvin5064/grpc-gateway/protobuf/database"
)

func (s *server) SetDatabaseInfo(ctx context.Context, request *database.DataEntry) (*database.DataEntry, error) {
	fmt.Println("Processing SetDatabaseInfo for request", request)
	s.database[request.Key] = request.Value
	return request, nil
}

func (s *server) GetDatabaseInfo(ctx context.Context, request *database.DataEntry) (*database.DataEntry, error) {
	fmt.Println("Processing GetDatabaseInfo for request", request)
	value, ok := s.database[request.Key]
	if ok {
		request.Value = value
	}
	return request, nil
}
