package main

import (
	"client/proto"
	context "context"
	fmt "fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	grpc "google.golang.org/grpc"
)

var TEST_PORT string
var SERVER_PORT string
var PROJECT_ID string
var TOPIC_ID string
var SUB_ID string

func main() {
	err := readEnv()
	if err != nil {
		log.Fatalf("Error in read env: %v", err)
	}

	connection, err := grpc.Dial("localhost:"+SERVER_PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	clientOperation := proto.NewOperationServiceClient(connection)
	var op string
	var number1, number2 float64

	for true {
		number := 7
		fmt.Println("0 - Exit\n1 - Go to Calculator")
		fmt.Scanf("%d\n", &number)
		if number == 0 {
			break
		}
		if number == 1 {
			fmt.Println("Write the Operation: (1 + 2) (3 - 8)")
			fmt.Scanf("%g %s %g\n", &number1, &op, &number2)
			err := operation(number1, number2, op, clientOperation)
			if err != nil {
				fmt.Println("The correct text to executed one operation is : \n(1 + 2), (3 - 8), (20 - 8), (15 / 5), (20 * 2)")
			}
		}
		number1 = 0
		number2 = 0
		op = ""
		time.Sleep(2 * time.Second)
	}

}

func operation(number1 float64, number2 float64, operation string, client proto.OperationServiceClient) error {
	if operation == "+" {
		sum(number1, number2, operation, client)
		return nil
	}
	if operation == "-" {
		subtraction(number1, number2, operation, client)
		return nil
	}
	if operation == "/" {
		division(number1, number2, operation, client)
		return nil
	}
	if operation == "*" {
		multiplication(number1, number2, operation, client)
		return nil
	}

	return fmt.Errorf("Error NotFound the operation")

}

func sum(number1 float64, number2 float64, operation string, client proto.OperationServiceClient) {
	request := &proto.SumRequest{
		Operation: &proto.Operation{
			Number1: number1,
			Number2: number2,
		},
	}

	res, err := client.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Error in execution: %v", err)
	}

	log.Println(res)
}

func division(number1 float64, number2 float64, operation string, client proto.OperationServiceClient) {
	request := &proto.DivisionRequest{
		Operation: &proto.Operation{
			Number1: number1,
			Number2: number2,
		},
	}

	res, err := client.Division(context.Background(), request)
	if err != nil {
		log.Fatalf("Error in execution: %v", err)
	}

	log.Println(res)
}

func multiplication(number1 float64, number2 float64, operation string, client proto.OperationServiceClient) {
	request := &proto.MultiplicationRequest{
		Operation: &proto.Operation{
			Number1: number1,
			Number2: number2,
		},
	}

	res, err := client.Multiplication(context.Background(), request)
	if err != nil {
		log.Fatalf("Error in execution: %v", err)
	}

	log.Println(res)
}

func subtraction(number1 float64, number2 float64, operation string, client proto.OperationServiceClient) {
	request := &proto.SubtractionRequest{
		Operation: &proto.Operation{
			Number1: number1,
			Number2: number2,
		},
	}

	res, err := client.Subtraction(context.Background(), request)
	if err != nil {
		log.Fatalf("Error in execution: %v", err)
	}

	log.Println(res)
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
