package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://patsore:<pass>@cluster0.dx5co.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	messagesCollection := client.Database("messages").Collection("messages")
	cursor, err := messagesCollection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &messages); err != nil {
		panic(err)
	}
	router := gin.Default()

	router.GET("/messages", getMessage)
	router.POST("/messages", postMessage)
	router.Run("localhost:8080")
}

func postMessage(c *gin.Context) {
	var receivedMessage messagej
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&receivedMessage); err != nil {
		return
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://patsore:<pass>6@cluster0.dx5co.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	messagesCollection := client.Database("messages").Collection("messages")

	// Add the new album to the slice.
	_, error := messagesCollection.InsertOne(ctx, receivedMessage)

	if error != nil {
		panic(error)
	}

	c.IndentedJSON(http.StatusCreated, receivedMessage)
}

var messages []message

type messagej struct {
	ID        string `json:"_id"`
	Sender    string `json:"sender"`
	Body      string `json:"body"`
	TimeStamp string `json:"timeStamp"`
}

type message struct {
	ID        string `bson:"_id"`
	Sender    string `bson:"sender"`
	Body      string `bson:"body"`
	TimeStamp string `bson:"timeStamp"`
}

func getMessage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, messages)
}
