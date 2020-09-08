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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var TEST_PORT string
var SERVER_PORT string
var PROJECT_ID string
var TOPIC_ID string
var SUB_ID string
var DB_HOST string
var DB_PORT string
var DB_NAME string
var DB_USER string
var DB_PASS string

/*
var Db *mongo.Client
var OperationDb *mongo.Collection
var MongoContext context.Context
*/
type OperationModel struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Number1     float64            `bson:"number1"`
	Operation   string             `bson:"operation"`
	Number2     float64            `bson:"number2"`
	Result      float64            `bson:"result"`
	DateCreated string             `bson:"dateCreated"`
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

	/*
		MongoContext = context.Background()

		Db, err = mongo.Connect(MongoContext,
			options.Client().ApplyURI("mongodb://"+DB_USER+":"+DB_PASS+"@"+DB_HOST+":"+DB_PORT))
		if err != nil {
			log.Fatal(err)
		}

		err = Db.Ping(MongoContext, nil)
		if err != nil {
			fmt.Println("debug")
			log.Fatal(err)
		}
		OperationDb = Db.Database("calculator-go").Collection("operations")

		fmt.Println("Server successfully started on port:" + SERVER_PORT)
	*/
	err = sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		fmt.Print("Got message: ")
		str := m.Data

		file, err := os.OpenFile("operations.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println("Error in Receive: %v", err)
		}

		log.SetOutput(file)

		op := extractOperation(string(str))
		log.Println(op)
		fmt.Println(op)

		/*r, err := createOperation(op)
		if err != nil {
			fmt.Println("Error in createOperation: %v", err)
		}
		fmt.Println(r)
		*/
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
		Number1:   n1,
		Number2:   n2,
		Result:    result,
		Operation: strOp,
	}
}

/*
func createOperation(op OperationModel) (string, error) {
	dataDb := OperationModel{
		Number1:     op.Number1,
		Number2:     op.Number2,
		Result:      op.Result,
		Operation:   op.Operation,
		DateCreated: time.Now().String(),
	}

	result, err := OperationDb.InsertOne(MongoContext, dataDb)
	if err != nil {
		return "", err
	}

	idResult := result.InsertedID.(primitive.ObjectID)

	return idResult.Hex(), nil
}
*/
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
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")

	return nil
}
