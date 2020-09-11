package server

import (
	"context"
	"fmt"
	"net"
	"query/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

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

		// create a value into which the single document can be decoded
		var elem proto.OperationDb
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, &elem)
	}

	response := &proto.ListAllResponse{
		Operations: results,
	}

	return response, nil

}

func NewServer() error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	proto.RegisterOperationDbServiceServer(grpcServer, &OperationDbServiceServer{})

	MongoContext = context.Background()
	Db, err = mongo.Connect(MongoContext,
		options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		return err
	}

	err = Db.Ping(MongoContext, nil)
	if err != nil {
		return err
	}
	OperationCollection = Db.Database("calculator-go").Collection("operations")

	fmt.Println("Server successfully started on port:50051")

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil

}
