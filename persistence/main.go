package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"persistence/proto"

	"cloud.google.com/go/pubsub"
	"github.com/golang/protobuf/ptypes"
	"github.com/joho/godotenv"
)

var TEST_PORT string
var SERVER_PORT string
var PROJECT_ID string
var TOPIC_ID string
var SUB_ID string

type DbOperationService struct{}

func (s *DbOperationService) CreateOperation(ctx context.Context, req *proto.CreateOperationRequest) (*proto.CreateOperationResponse, error) {
	return &proto.CreateOperationResponse{
		Id: "1",
	}, nil
}
func (s *DbOperationService) DeleteOperation(ctx context.Context, req *proto.DeleteOperationRequest) (*proto.DeleteOperationResponse, error) {
	return &proto.DeleteOperationResponse{
		Success: true,
	}, nil
}

func (s *DbOperationService) ReadOperation(ctx context.Context, req *proto.ReadOperationRequest) (*proto.ReadOperationResponse, error) {

	return &proto.ReadOperationResponse{
		DbOperation: &proto.DbOperation{
			Id:          "id",
			Number1:     1,
			Number2:     2,
			Result:      3,
			Operation:   "+",
			DateCreated: ptypes.TimestampNow(),
		},
	}, nil
}

func (s *DbOperationService) ListOperation(ctx context.Context, req *proto.ListOperationRequest) (*proto.ListOperationResponse, error) {
	return &proto.ListOperationResponse{
		DbOperations: []*proto.DbOperation{
			{Id: "id",
				Number1:     1,
				Number2:     2,
				Result:      3,
				Operation:   "+",
				DateCreated: ptypes.TimestampNow(),
			},
		},
	}, nil

}

func main() {
	ctx := context.Background()

	err := readEnv()
	if err != nil {
		log.Fatalf("Error in read env: %v", err)
	}

	clientPubSub, err := pubsub.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Fatalf("Error in connection to clientPubSub: %v", err)
	}

	topic, err := clientPubSub.CreateTopic(ctx, TOPIC_ID)
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

	err = sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message: %s", m.Data)
		m.Ack()
	})
	if err != nil {
		fmt.Println("Error in Receive: %v", err)
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
