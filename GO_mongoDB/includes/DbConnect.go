package includes

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://RajendroSau:rajandrO21@cluster0.nwqhn1l.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println("Database Name", databases)
	return client
}

// Client instance
var DB *mongo.Client = Connect()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(DB_NAME).Collection(collectionName)
	return collection
}
