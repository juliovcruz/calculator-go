package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"operation/proto"
	"os"
	"strconv"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var TEST_PORT string
var SERVER_PORT string
var PROJECT_ID string
var TOPIC_ID string
var SUB_ID string
var topic *pubsub.Topic

type OperationServiceServer struct{}

func (s *OperationServiceServer) Division(ctx context.Context, req *proto.DivisionRequest) (*proto.DivisionResponse, error) {
	data := req.GetOperation()
	num1 := data.GetNumber1()
	num2 := data.GetNumber2()

	result := num1 / num2

	publishMessage(ctx, topic, strconv.FormatFloat(num1, 'g', -1, 64)+"/"+strconv.FormatFloat(num2, 'g', -1, 64)+"="+strconv.FormatFloat(result, 'g', -1, 64))

	return &proto.DivisionResponse{Result: result}, nil
}

func (s *OperationServiceServer) Multiplication(ctx context.Context, req *proto.MultiplicationRequest) (*proto.MultiplicationResponse, error) {
	data := req.GetOperation()
	num1 := data.GetNumber1()
	num2 := data.GetNumber2()

	result := num1 * num2

	publishMessage(ctx, topic, strconv.FormatFloat(num1, 'g', -1, 64)+"*"+strconv.FormatFloat(num2, 'g', -1, 64)+"="+strconv.FormatFloat(result, 'g', -1, 64))

	return &proto.MultiplicationResponse{Result: result}, nil
}

func (s *OperationServiceServer) Sum(ctx context.Context, req *proto.SumRequest) (*proto.SumResponse, error) {
	data := req.GetOperation()
	num1 := data.GetNumber1()
	num2 := data.GetNumber2()

	result := num1 + num2

	publishMessage(ctx, topic, strconv.FormatFloat(num1, 'g', -1, 64)+"+"+strconv.FormatFloat(num2, 'g', -1, 64)+"="+strconv.FormatFloat(result, 'g', -1, 64))

	return &proto.SumResponse{Result: result}, nil
}

func (s *OperationServiceServer) Subtraction(ctx context.Context, req *proto.SubtractionRequest) (*proto.SubtractionResponse, error) {
	data := req.GetOperation()
	num1 := data.GetNumber1()
	num2 := data.GetNumber2()
	result := num1 - num2

	publishMessage(ctx, topic, strconv.FormatFloat(num1, 'g', -1, 64)+"-"+strconv.FormatFloat(num2, 'g', -1, 64)+"="+strconv.FormatFloat(result, 'g', -1, 64))

	return &proto.SubtractionResponse{Result: result}, nil
}

func publishMessage(ctx context.Context, topic *pubsub.Topic, msg string) {
	topic.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
}

func main() {
	ctx := context.Background()

	err := readEnv()
	if err != nil {
		log.Fatalf("Error in read env: %v", err)
	}

	listener, err := net.Listen("tcp", ":"+SERVER_PORT)
	if err != nil {
		log.Fatalf("Error in listen port:" + SERVER_PORT)
	}

	clientPubSub, err := pubsub.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Fatalf("Error in connection to clientPubSub: %v", err)
	}

	topic, err = clientPubSub.CreateTopic(ctx, TOPIC_ID)
	if err != nil {
		topic = clientPubSub.Topic(TOPIC_ID)
		fmt.Println("Fail in createTopic: %v", err)
	}
	defer topic.Stop()

	sub, err := clientPubSub.CreateSubscription(ctx, SUB_ID, pubsub.SubscriptionConfig{
		Topic: topic,
	})
	if err != nil {
		sub = clientPubSub.Subscription(SUB_ID)
		fmt.Println("Fail in createSubscription: %v", err)
	}
	fmt.Println(sub)

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
	PROJECT_ID = os.Getenv("PROJECT_ID")
	TOPIC_ID = os.Getenv("TOPIC_ID")
	SUB_ID = os.Getenv("SUB_ID")

	return nil
}
