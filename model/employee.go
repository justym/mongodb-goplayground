package model

type Employee struct {
	Name string   `bson:"name"`
	Year string   `bson:"year"`
	Jobs []string `bson:"jobs"`
}
