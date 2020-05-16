package dbutil

import (
	"context"
	"log"

	"github.com/justym/mongodb-goplayground/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func NewClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	return client, nil
}

func NewCollection(client *mongo.Client) (*mongo.Collection, error) {
	log.Println("Connected to MongoDB successfully")
	collection := client.Database("employeeDB").Collection("employee")

	return collection, nil
}

func InsertOne(e *model.Employee, c *mongo.Collection) error {
	_, err := c.InsertOne(context.TODO(), *e)
	if err != nil {
		return err
	}

	return nil
}

func FindOne(e *model.Employee, collection *mongo.Collection) error {
	queryResult := &model.Employee{}
	filter := bson.M{"name": bson.M{"$eq": e.Name}}
	res := collection.FindOne(context.TODO(), filter)
	if err := res.Decode(queryResult); err != nil {
		return err
	}

	log.Printf("Employee: %v, %v, %v", queryResult.Name, queryResult.Year, queryResult.Jobs)
	return nil
}

func Disconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	log.Println("MongoDB client disconnected successfully")
}
