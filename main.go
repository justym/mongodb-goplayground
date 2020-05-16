package main

import (
	"fmt"
	"log"

	"github.com/justym/mongodb-goplayground/model"
	"github.com/justym/mongodb-goplayground/dbutil"
)

var mockDoc = model.Employee{
	Name: "Mock",
	Year: "2018",
	Jobs: []string{"product_manager"},
}

func main() {
	client, err := dbutil.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer dbutil.Disconnect(client)

	collection, err := dbutil.NewCollection(client)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbutil.InsertOne(&mockDoc, collection); err != nil {
		err = fmt.Errorf("[insertOne] %+v",err)
		log.Fatal(err)
	}

	if err := dbutil.FindOne(&mockDoc, collection); err != nil {
		err = fmt.Errorf("[findOne] %+v", err)
		log.Fatal(err)
	}
}

