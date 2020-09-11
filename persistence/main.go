package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var PROJECT_ID, TOPIC_ID, SUB_ID, DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASS string

var Db *mongo.Client
var OperationDb *mongo.Collection
var MongoContext context.Context

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

	MongoContext = context.Background()
	//options.Client().ApplyURI("mongodb://"+DB_USER+":"+DB_PASS+"@"+DB_HOST+":"+DB_PORT))
	Db, err = mongo.Connect(MongoContext,
		options.Client().ApplyURI("mongodb://"+DB_HOST+":"+DB_PORT+"/"))
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping(MongoContext, nil)
	if err != nil {
		log.Fatal(err)
	}
	OperationDb = Db.Database("calculator-go").Collection("operations")

	fmt.Println("Server successfully started on port:" + SERVER_PORT)

	err = sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		var op OperationModel
		err := json.Unmarshal(m.Data, &op)
		if err != nil {
			fmt.Println(err)
		}

		resultOp, err := createOperation(op)
		if err != nil {
			fmt.Println("Error in createOperation: %v", err)
		}
		fmt.Println(resultOp)

		m.Ack()
	})
	if err != nil {
		fmt.Println("Error in Receive: %v", err)
	}

}

func createOperation(op OperationModel) (*OperationModel, error) {
	dataDb := OperationModel{
		Number1:     op.Number1,
		Number2:     op.Number2,
		Result:      op.Result,
		Operation:   op.Operation,
		DateCreated: time.Now().String(),
	}

	result, err := OperationDb.InsertOne(MongoContext, dataDb)
	if err != nil {
		return nil, err
	}

	idResult := result.InsertedID.(primitive.ObjectID)
	dataDb.Id = idResult

	return &dataDb, nil
}

func readEnv() error {
	err := godotenv.Load("../.env")
	if err != nil {
		return err
	}
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
