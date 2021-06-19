package main

import (
	"context"
	"gitabza-go/mongodb"
	"log"
)

func main() {
	ctx := context.Background()

	_, err := mongodb.Connect(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("error connectiing to mongodb: %v\n", err)
	}
	defer func() {
		if err := mongodb.Disconnect(ctx); err != nil {
			log.Fatalf("error disconnecting from mongodb: %v\n", err)
		}
		log.Println("disconnected from mongodb")
	}()
}
