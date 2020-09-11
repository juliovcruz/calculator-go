package main

import (
	"fmt"
	"query/server"
)

func main() {
	err := server.NewServer()
	if err != nil {
		fmt.Println(err)
	}
}
