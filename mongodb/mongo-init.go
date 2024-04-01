package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func MongoInit() *mongo.Client {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	return client
}
func Ping() string {
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return "Failed to Connect to DB"
	}
	return "Connected to DB successfully"
}

func GetAllCollections(db string) []string {
	db_name := client.Database(db)
	fmt.Println(db_name.Name())
	True := true
	options := options.ListCollectionsOptions{
		NameOnly: &True,
	}
	cursor, _ := db_name.ListCollectionNames(context.Background(), bson.M{}, &options)
	fmt.Println("Collections:", cursor)
	return cursor

}
