package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/juliovcruz/calculator-go/query/proto"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var QUERY_PORT, DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASS string

var Db *mongo.Client
var OperationCollection *mongo.Collection
var MongoContext context.Context

var StringToEnum = map[string]proto.EnumOp{
	"NULL":           proto.EnumOp_NULL,
	"DIVISION":       proto.EnumOp_DIVISION,
	"MULTIPLICATION": proto.EnumOp_MULTIPLICATION,
	"SUBTRACTION":    proto.EnumOp_SUBTRACTION,
	"SUM":            proto.EnumOp_SUM,
}

type OperationModel struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Number1     float64            `bson:"number1"`
	Operation   string             `bson:"operation"`
	Number2     float64            `bson:"number2"`
	Result      float64            `bson:"result"`
	DateCreated string             `bson:"dateCreated"`
}

type OperationDbServiceServer struct{}

func (s *OperationDbServiceServer) FindById(ctx context.Context, req *proto.FindByIdRequest) (*proto.FindByIdResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, err
	}

	result := OperationCollection.FindOne(ctx, bson.M{"_id": id})
	data := OperationModel{}
	if err := result.Decode(&data); err != nil {
		return nil, err
	}

	response := &proto.FindByIdResponse{
		OperationDb: &proto.OperationDb{
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

func (s *OperationDbServiceServer) FindByOp(ctx context.Context, req *proto.FindByOpRequest) (*proto.FindByOpResponse, error) {

	cursor, err := OperationCollection.Find(ctx, bson.M{"operation": req.GetOp().String()})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*proto.OperationDb

	for cursor.Next(context.TODO()) {

		var elem OperationModel
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		elemResult := proto.OperationDb{
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

func (s *OperationDbServiceServer) ListAll(ctx context.Context, req *proto.ListAllRequest) (*proto.ListAllResponse, error) {

	cursor, err := OperationCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*proto.OperationDb

	for cursor.Next(context.TODO()) {

		var elem OperationModel
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		elemResult := proto.OperationDb{
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

func NewServer() error {

	err := readEnv()
	if err != nil {
		log.Fatalf("Error in read env: %v", err)
	}

	listener, err := net.Listen("tcp", ":"+QUERY_PORT)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	proto.RegisterOperationDbServiceServer(grpcServer, &OperationDbServiceServer{})

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

	fmt.Println("Server successfully started on port:" + QUERY_PORT)

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

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	QUERY_PORT = os.Getenv("QUERY_PORT")

	return nil
}
