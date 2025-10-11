package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client{
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	// conneting to mongodb
	client,err := mongo.Connect(ctx,clientOptions)
	if err != nil{
		log.Fatal(err)
	}

	err = client.Ping(ctx,nil)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to mongodb")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection()
