package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/juliovcruz/calculator-go/operation/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var SERVER_PORT, PROJECT_ID, TOPIC_ID, SUB_ID string
var QUERY_PORT, DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASS string

var Db *mongo.Client
var OperationCollection *mongo.Collection
var MongoContext context.Context

type ServerOptions struct {
	Pubsub bool
}

type OperationServiceServer struct {
	topic *pubsub.Topic
}

type Operation struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Number1     float64            `bson:"number1"`
	Operation   string             `bson:"operation"`
	Number2     float64            `bson:"number2"`
	Result      float64            `bson:"result"`
	DateCreated string             `bson:"dateCreated"`
}

func (s *OperationServiceServer) Division(ctx context.Context, req *proto.DivisionRequest) (*proto.DivisionResponse, error) {
	num1 := req.GetNumber1()
	num2 := req.GetNumber2()

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
	num1 := req.GetNumber1()
	num2 := req.GetNumber2()

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
	num1 := req.GetNumber1()
	num2 := req.GetNumber2()

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
	num1 := req.GetNumber1()
	num2 := req.GetNumber2()

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

func (s *OperationServiceServer) FindById(ctx context.Context, req *proto.FindByIdRequest) (*proto.FindByIdResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, err
	}

	result := OperationCollection.FindOne(ctx, bson.M{"_id": id})
	data := Operation{}
	if err := result.Decode(&data); err != nil {
		return nil, err
	}

	response := &proto.FindByIdResponse{
		Operation: &proto.Operation{
			//Id:          id.Hex(),
			Number1:     data.Number1,
			Operation:   data.Operation,
			Number2:     data.Number2,
			Result:      data.Result,
			DateCreated: data.DateCreated,
		},
	}

	return response, nil

}

func (s *OperationServiceServer) FindByOp(ctx context.Context, req *proto.FindByOpRequest) (*proto.FindByOpResponse, error) {

	cursor, err := OperationCollection.Find(ctx, bson.M{"operation": req.GetOp().String()})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*proto.Operation

	for cursor.Next(context.TODO()) {

		var elem Operation
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		elemResult := proto.Operation{
			Id:          elem.Id.Hex(),
			Number1:     elem.Number1,
			Operation:   elem.Operation,
			Number2:     elem.Number2,
			Result:      elem.Result,
			DateCreated: elem.DateCreated,
		}

		results = append(results, &elemResult)
	}

	response := &proto.FindByOpResponse{
		Operations: results,
	}

	return response, nil

}

func (s *OperationServiceServer) ListAll(ctx context.Context, req *proto.ListAllRequest) (*proto.ListAllResponse, error) {

	cursor, err := OperationCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*proto.Operation

	for cursor.Next(context.TODO()) {

		var elem Operation
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		elemResult := proto.Operation{
			Id:          elem.Id.Hex(),
			Number1:     elem.Number1,
			Operation:   elem.Operation,
			Number2:     elem.Number2,
			Result:      elem.Result,
			DateCreated: elem.DateCreated,
		}

		results = append(results, &elemResult)
	}

	response := &proto.ListAllResponse{
		Operations: results,
	}

	return response, nil

}

func NewServer(ctx context.Context, serverOptions ServerOptions) error {

	err := readEnv()
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", ":"+SERVER_PORT)
	if err != nil {
		return err
	}

	serverOperation := &OperationServiceServer{}

	if serverOptions.Pubsub {
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

	MongoContext = context.Background()

	Db, err = mongo.Connect(MongoContext,
		options.Client().ApplyURI("mongodb://"+DB_HOST+":"+DB_PORT+"/"))
	if err != nil {
		return err
	}

	err = Db.Ping(MongoContext, nil)
	if err != nil {
		return err
	}
	OperationCollection = Db.Database("calculator-go").Collection("operations")

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
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")

	return nil
}
