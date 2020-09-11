package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/juliovcruz/calculator-go/operation/proto"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var SERVER_PORT, PROJECT_ID, TOPIC_ID, SUB_ID string

type ServerOptions struct {
	Pubsub bool
}

type OperationServiceServer struct {
	topic *pubsub.Topic
}

type Operation struct {
	Number1   float64
	Number2   float64
	Result    float64
	Operation string
}

func (s *OperationServiceServer) Division(ctx context.Context, req *proto.DivisionRequest) (*proto.DivisionResponse, error) {
	data := req.GetOperation()
	num1 := data.GetNumber1()
	num2 := data.GetNumber2()

	result := num1 / num2

	if num1 == 0 || num2 == 0 {
		result = 0
	}

	opData := Operation{Number1: num1, Number2: num2, Result: result, Operation: "DIVISION"}

	if s.topic != nil {
		byteJson, err := json.Marshal(opData)
		if err != nil {
			return nil, err
		}
		s.topic.Publish(ctx, &pubsub.Message{
			Data: []byte(byteJson),
		})
	}

	return &proto.DivisionResponse{Result: result}, nil
}

func (s *OperationServiceServer) Multiplication(ctx context.Context, req *proto.MultiplicationRequest) (*proto.MultiplicationResponse, error) {
	data := req.GetOperation()
	num1 := data.GetNumber1()
	num2 := data.GetNumber2()

	result := num1 * num2

	opData := Operation{Number1: num1, Number2: num2, Result: result, Operation: "MULTIPLICATION"}

	if s.topic != nil {
		byteJson, err := json.Marshal(opData)
		if err != nil {
			return nil, err
		}
		s.topic.Publish(ctx, &pubsub.Message{
			Data: []byte(byteJson),
		})
	}

	return &proto.MultiplicationResponse{Result: result}, nil
}

func (s *OperationServiceServer) Sum(ctx context.Context, req *proto.SumRequest) (*proto.SumResponse, error) {
	data := req.GetOperation()
	num1 := data.GetNumber1()
	num2 := data.GetNumber2()

	result := num1 + num2

	opData := Operation{Number1: num1, Number2: num2, Result: result, Operation: "SUM"}

	if s.topic != nil {
		byteJson, err := json.Marshal(opData)
		if err != nil {
			return nil, err
		}
		s.topic.Publish(ctx, &pubsub.Message{
			Data: []byte(byteJson),
		})
	}

	return &proto.SumResponse{Result: result}, nil
}

func (s *OperationServiceServer) Subtraction(ctx context.Context, req *proto.SubtractionRequest) (*proto.SubtractionResponse, error) {
	data := req.GetOperation()
	num1 := data.GetNumber1()
	num2 := data.GetNumber2()

	result := num1 - num2

	opData := Operation{Number1: num1, Number2: num2, Result: result, Operation: "SUBTRACTION"}

	if s.topic != nil {
		byteJson, err := json.Marshal(opData)
		if err != nil {
			return nil, err
		}
		s.topic.Publish(ctx, &pubsub.Message{
			Data: []byte(byteJson),
		})
	}

	return &proto.SubtractionResponse{Result: result}, nil
}

func NewServer(ctx context.Context, options ServerOptions) error {

	err := readEnv()
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", ":"+SERVER_PORT)
	if err != nil {
		return err
	}

	serverOperation := &OperationServiceServer{}

	if options.Pubsub {
		clientPubSub, err := pubsub.NewClient(ctx, PROJECT_ID)
		if err != nil {
			return err
		}
		serverOperation.topic, err = clientPubSub.CreateTopic(ctx, TOPIC_ID)
		if err != nil { // identificar erro
			serverOperation.topic = clientPubSub.Topic(TOPIC_ID)
			fmt.Println(err)
		}
		defer serverOperation.topic.Stop()

		sub, err := clientPubSub.CreateSubscription(ctx, SUB_ID, pubsub.SubscriptionConfig{
			Topic: serverOperation.topic,
		})
		if err != nil {
			sub = clientPubSub.Subscription(SUB_ID)
			fmt.Println(err)
		}
		fmt.Println(sub)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterOperationServiceServer(grpcServer, serverOperation)

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

func readEnv() error {
	err := godotenv.Load("../.env")
	if err != nil {
		return err
	}

	SERVER_PORT = os.Getenv("SERVER_PORT")
	PROJECT_ID = os.Getenv("PROJECT_ID")
	TOPIC_ID = os.Getenv("TOPIC_ID")
	SUB_ID = os.Getenv("SUB_ID")

	return nil
}
