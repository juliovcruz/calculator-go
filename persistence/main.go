package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

var TEST_PORT string
var SERVER_PORT string
var PROJECT_ID string
var TOPIC_ID string
var SUB_ID string

type OperationModel struct {
	number1   float64
	number2   float64
	operation string
	result    float64
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
		str := m.Data

		op := extractOperation(string(str))

		fmt.Println(op)

		m.Ack()
	})
	if err != nil {
		fmt.Println("Error in Receive: %v", err)
	}

}

func extractOperation(str string) OperationModel {
	var strN1, strN2, strOp, strResult string
	if strings.ContainsAny(str, "-") {
		strOp = "-"
		strN1 = str[0:strings.IndexAny(str, strOp)]
		strN2 = str[strings.IndexAny(str, strOp)+1 : strings.IndexAny(str, "=")]
		strResult = str[strings.IndexAny(str, "=")+1 : len(str)]
	}
	if strings.ContainsAny(str, "+") {
		strOp = "+"
		strN1 = str[0:strings.IndexAny(str, strOp)]
		strN2 = str[strings.IndexAny(str, strOp)+1 : strings.IndexAny(str, "=")]
		strResult = str[strings.IndexAny(str, "=")+1 : len(str)]
	}
	if strings.ContainsAny(str, "/") {
		strOp = "/"
		strN1 = str[0:strings.IndexAny(str, strOp)]
		strN2 = str[strings.IndexAny(str, strOp)+1 : strings.IndexAny(str, "=")]
		strResult = str[strings.IndexAny(str, "=")+1 : len(str)]
	}
	if strings.ContainsAny(str, "*") {
		strOp = "*"
		strN1 = str[0:strings.IndexAny(str, strOp)]
		strN2 = str[strings.IndexAny(str, strOp)+1 : strings.IndexAny(str, "=")]
		strResult = str[strings.IndexAny(str, "=")+1 : len(str)]
	}

	n1, _ := strconv.ParseFloat(strN1, 64)
	n2, _ := strconv.ParseFloat(strN2, 64)
	result, _ := strconv.ParseFloat(strResult, 64)

	return OperationModel{
		number1:   n1,
		number2:   n2,
		result:    result,
		operation: strOp,
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
