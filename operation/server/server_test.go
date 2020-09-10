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

func TestSum(t *testing.T) {
	s := OperationServiceServer{}

	mapTest := map[float64]OperationModel{
		10: OperationModel{
			number1: 5,
			number2: 5,
		},
		20: OperationModel{
			number1: 1,
			number2: 19,
		},
	}

	t.Run("Sum",
		func(t *testing.T) {
			for key, value := range mapTest {
				res, err := s.Sum(context.Background(),
					&proto.SumRequest{
						Operation: &proto.Operation{
							Number1: value.number1,
							Number2: value.number2,
						},
					},
				)
				if err != nil {
					t.Fatalf("Error in execution method %v:", err)
				}
				assert.Equal(t, key, res.Result)
			}
		},
	)

}

func TestDivision(t *testing.T) {
	s := OperationServiceServer{}

	mapTest := map[float64]OperationModel{
		20: OperationModel{
			number1: 40,
			number2: 2,
		},
		0: OperationModel{
			number1: 5,
			number2: 0,
		},
		30: OperationModel{
			number1: 60,
			number2: 2,
		},
	}

	t.Run("Division",
		func(t *testing.T) {
			for key, value := range mapTest {
				res, err := s.Division(context.Background(),
					&proto.DivisionRequest{
						Operation: &proto.Operation{
							Number1: value.number1,
							Number2: value.number2,
						},
					},
				)
				if err != nil {
					t.Fatalf("Error in execution method %v:", err)
				}
				assert.Equal(t, key, res.Result)
			}
		},
	)

}
