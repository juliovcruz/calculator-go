package main

import (
	"log"
	"net"
	"operation/proto"
	"os"
	"testing"

	"google.golang.org/grpc"
)

type OperationModel struct {
	number1 float64
	number2 float64
}

func Server() {
	err := readEnv()
	if err != nil {
		log.Fatalf("Error in readEnv: %v", err)
	}

	listener, err := net.Listen("tcp", ":"+TEST_PORT)
	if err != nil {
		log.Fatalf("Error in listen port: " + TEST_PORT)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterOperationServiceServer(grpcServer, &OperationServiceServer{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error in connection gRpc with listener")
	}
}

func TestMain(m *testing.M) {
	go Server()
	os.Exit(m.Run())
}

func TestSum(t *testing.T) {
	err := readEnv()
	if err != nil {
		log.Fatalf("Error in readEnv: %v", err)
	}

	conn, err := grpc.Dial("localhost:"+TEST_PORT, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Error in connection to ServerTest: %v", err)
	}
	defer conn.Close()
	server := proto.NewOperationServiceClient(conn)

	mapTest := map[float64]OperationModel{
		10: OperationModel{
			number1: 5,
			number2: 5,
		},
	}

}
