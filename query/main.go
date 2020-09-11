package main

import (
	"fmt"

	"github.com/juliovcruz/calculator-go/query/server"
)

func main() {
	err := server.NewServer()
	if err != nil {
		fmt.Println(err)
	}
}
