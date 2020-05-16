package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Employee struct {
	Name string   `bson:"name"`
	Year string   `bson:"year"`
	Jobs []string `bson:"jobs"`
}

var mockDoc = Employee{
	Name: "Mock",
	Year: "2018",
	Jobs: []string{"product_manager"},
}

func main() {
	client, err := newClient()
	if err != nil {
		log.Fatal(err)
	}
	defer disconnect(client)

	collection, err := newCollection(client)
	if err != nil {
		log.Fatal(err)
	}

	if err := insertOne(&mockDoc, collection); err != nil {
		err = fmt.Errorf("[insertOne] %+v",err)
		log.Fatal(err)
	}

	if err := findOne(&mockDoc, collection); err != nil {
		err = fmt.Errorf("[findOne] %+v", err)
		log.Fatal(err)
	}
}

func newClient() (*mongo.Client, error) {
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

func newCollection(client *mongo.Client) (*mongo.Collection, error) {
	log.Println("Connected to MongoDB successfully")
	collection := client.Database("employeeDB").Collection("employee")

	return collection, nil
}

func insertOne(e *Employee, c *mongo.Collection) error {
	_, err := c.InsertOne(context.TODO(), *e)
	if err != nil {
		return err
	}

	return nil
}

func disconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	log.Println("MongoDB client disconnected successfully")
}

func findOne(e *Employee, collection *mongo.Collection) error {
	queryResult := &Employee{}
	filter := bson.M{"name": bson.M{"$eq": e.Name}}
	res := collection.FindOne(context.TODO(),filter)
	if err := res.Decode(queryResult); err != nil {
		return err
	}

	log.Printf("Employee: %v, %v, %v", queryResult.Name, queryResult.Year, queryResult.Jobs)
	return nil
}