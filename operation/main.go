package main

import (
	"context"
	"log"
	"net"
	"operation/proto"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var TEST_PORT string
var SERVER_PORT string

type OperationServiceServer struct{}

func (s *OperationServiceServer) Division(ctx context.Context, req *proto.DivisionRequest) (*proto.DivisionResponse, error) {
	data := req.GetOperation()

	result := data.GetNumber1() / data.GetNumber2()

	return &proto.DivisionResponse{Result: result}, nil
}

func (s *OperationServiceServer) Multiplication(ctx context.Context, req *proto.MultiplicationRequest) (*proto.MultiplicationResponse, error) {
	data := req.GetOperation()

	result := data.GetNumber1() * data.GetNumber2()

	return &proto.MultiplicationResponse{Result: result}, nil
}

func (s *OperationServiceServer) Sum(ctx context.Context, req *proto.SumRequest) (*proto.SumResponse, error) {
	data := req.GetOperation()

	result := data.GetNumber1() + data.GetNumber2()

	return &proto.SumResponse{Result: result}, nil
}

func (s *OperationServiceServer) Subtraction(ctx context.Context, req *proto.SubtractionRequest) (*proto.SubtractionResponse, error) {
	return &proto.SubtractionResponse{Result: 1}, nil
}

func main() {
	err := readEnv()
	if err != nil {
		log.Fatalf("Error in read env: %v", err)
	}

	listener, err := net.Listen("tcp", ":"+SERVER_PORT)
	if err != nil {
		log.Fatalf("Error in listen port:" + SERVER_PORT)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterOperationServiceServer(grpcServer, &OperationServiceServer{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error in connection gRpc with listener")
	}

}

func readEnv() error {
	err := godotenv.Load("../.env")
	if err != nil {
		return err
	}
	TEST_PORT = os.Getenv("TEST_PORT")
	SERVER_PORT = os.Getenv("SERVER_PORT")
	return nil
}
