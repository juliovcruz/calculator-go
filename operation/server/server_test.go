package server

import (
	"context"
	"operation/proto"
	"testing"

	"github.com/stretchr/testify/assert"
)

type OperationModel struct {
	number1 float64
	number2 float64
}

type InterfaceTest struct {
	op     OperationModel
	result float64
}

func TestSum(t *testing.T) {
	s := OperationServiceServer{}

	sliceTest := []InterfaceTest{
		InterfaceTest{
			op: OperationModel{
				number1: 10,
				number2: 2,
			},
			result: 12,
		},
	}

	t.Run("Sum",
		func(t *testing.T) {
			for _, value := range sliceTest {
				res, err := s.Sum(context.Background(),
					&proto.SumRequest{
						Operation: &proto.Operation{
							Number1: value.op.number1,
							Number2: value.op.number2,
						},
					},
				)
				if err != nil {
					t.Fatalf("Error in execution method %v:", err)
				}
				assert.Equal(t, value.result, res.Result)
			}
		},
	)

}

func TestDivision(t *testing.T) {
	s := OperationServiceServer{}

	sliceTest := []InterfaceTest{
		InterfaceTest{
			op: OperationModel{
				number1: 20,
				number2: 2,
			},
			result: 10,
		},
		InterfaceTest{
			op: OperationModel{
				number1: 10,
				number2: 5,
			},
			result: 2,
		},
		InterfaceTest{
			op: OperationModel{
				number1: 10,
				number2: 0,
			},
			result: 0,
		},
	}

	t.Run("Division",
		func(t *testing.T) {
			for _, value := range sliceTest {
				res, err := s.Division(context.Background(),
					&proto.DivisionRequest{
						Operation: &proto.Operation{
							Number1: value.op.number1,
							Number2: value.op.number2,
						},
					},
				)
				if err != nil {
					t.Fatalf("Error in execution method %v:", err)
				}
				assert.Equal(t, value.result, res.Result)
			}
		},
	)

}

func TestMultiplication(t *testing.T) {
	s := OperationServiceServer{}

	sliceTest := []InterfaceTest{
		InterfaceTest{
			op: OperationModel{
				number1: 20,
				number2: 2,
			},
			result: 40,
		},
		InterfaceTest{
			op: OperationModel{
				number1: 10,
				number2: 5,
			},
			result: 50,
		},
		InterfaceTest{
			op: OperationModel{
				number1: 10,
				number2: 0,
			},
			result: 0,
		},
	}

	t.Run("Multiplication",
		func(t *testing.T) {
			for _, value := range sliceTest {
				res, err := s.Multiplication(context.Background(),
					&proto.MultiplicationRequest{
						Operation: &proto.Operation{
							Number1: value.op.number1,
							Number2: value.op.number2,
						},
					},
				)
				if err != nil {
					t.Fatalf("Error in execution method %v:", err)
				}
				assert.Equal(t, value.result, res.Result)
			}
		},
	)

}

func TestSubtraction(t *testing.T) {
	s := OperationServiceServer{}

	sliceTest := []InterfaceTest{
		InterfaceTest{
			op: OperationModel{
				number1: 20,
				number2: 2,
			},
			result: 18,
		},
		InterfaceTest{
			op: OperationModel{
				number1: 10,
				number2: 5,
			},
			result: 5,
		},
		InterfaceTest{
			op: OperationModel{
				number1: 10,
				number2: 0,
			},
			result: 10,
		},
	}

	t.Run("Subtraction",
		func(t *testing.T) {
			for _, value := range sliceTest {
				res, err := s.Subtraction(context.Background(),
					&proto.SubtractionRequest{
						Operation: &proto.Operation{
							Number1: value.op.number1,
							Number2: value.op.number2,
						},
					},
				)
				if err != nil {
					t.Fatalf("Error in execution method %v:", err)
				}
				assert.Equal(t, value.result, res.Result)
			}
		},
	)

}
