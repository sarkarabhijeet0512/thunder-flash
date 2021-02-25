package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	// Import the JSON

	// encoding package

	// Official 'mongo-go-driver' packages
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoFields type struct
type MongoFields struct {
	//Key string `json:"key,omitempty"`
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Popularity float64            `json:"popularity,omitempty" bson:"popularity,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Director   string             `json:"director,omitempty" bson:"director,omitempty"`
	Imdbscore  float64            `json:"imdbscore,omitempty" bson:"imdbscore,omitempty"`
	Genre      []string           `json:"genre" bson:"genre,omitempty"`
}

func main() {
	fmt.Println("server is starting.......")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("mongo.Connect() ERROR: %v", err)
	}
	collection := client.Database("JSON_docs").Collection("JSON Collection")
	docsPath, _ := filepath.Abs("imdb.json") //Abs gives absolute path of the file
	byteValues, err := ioutil.ReadFile(docsPath)
	if err != nil {
		fmt.Println("ioutil.ReadFile ERROR:", err)
	} else {
		// fmt.Println("ioutil.ReadFile byteValues TYPE:", reflect.TypeOf(byteValues))
		// fmt.Println("byteValues:", byteValues, "n")
		// fmt.Println("byteValues:", string(byteValues))
		var doc []interface{}
		err = json.Unmarshal(byteValues, &doc)
		result, insertErr := collection.InsertMany(ctx, doc)
		if insertErr != nil {
			fmt.Println("InsertOne ERROR:", insertErr)
		} else {
			fmt.Println("InsertOne() API result:", result)
		}
	}
}
