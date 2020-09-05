package main

import (
	"context"
	"persistence/proto"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DbOperationService struct{}

var Db *mongo.Client
var OperationDb *mongo.Collection
var MongoContext context.Context

type OperationModel struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Number1     float64            `bson:"number1"`
	Number2     float64            `bson:"number2"`
	Result      float64            `bson:"result"`
	Operation   string             `bson:"operation"`
	DateCreated string             `bson:"dateCreated"`
}

func (s *DbOperationService) CreateOperation(ctx context.Context, req *proto.CreateOperationRequest) (*proto.CreateOperationResponse, error) {
	data := req.GetDbOperation()

	dataDb := OperationModel{
		Number1:     data.GetNumber1(),
		Number2:     data.GetNumber2(),
		Result:      data.GetResult(),
		Operation:   data.GetOperation(),
		DateCreated: time.Now().String(),
	}

	result, err := OperationDb.InsertOne(MongoContext, dataDb)
	if err != nil {
		return nil, err
	}

	idResult := result.InsertedID.(primitive.ObjectID)

	return &proto.CreateOperationResponse{
		Id: idResult.Hex(),
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
			DateCreated: time.Now().String(),
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
				DateCreated: time.Now().String(),
			},
		},
	}, nil

}
