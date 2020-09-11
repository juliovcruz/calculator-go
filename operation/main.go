package main

import (
	"context"
	"log"

	"github.com/juliovcruz/calculator-go/operation/server"
)

func main() {
	ctx := context.Background()

	err := server.NewServer(ctx, server.ServerOptions{
		Pubsub: true,
	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

}
