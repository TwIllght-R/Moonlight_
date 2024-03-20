package main

import (
	"Moonlight_/core"
	"Moonlight_/handler"
	"Moonlight_/repo"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client := initDatabase()
	initTimeZone()

	userCollection := client.Database("lala").Collection("users")       //users
	storyCollection := client.Database("lala").Collection("storyies")   //tasks
	commentCollection := client.Database("lala").Collection("comments") //comments

	//user
	userRepo := repo.NewUserRepo(userCollection)
	userCore := core.NewUserCore(userRepo)
	//task
	storyRepo := repo.NewStoryRepo(storyCollection)
	storyCore := core.NewStoryCore(storyRepo)
	//comment
	commentRepo := repo.NewCommentRepo(commentCollection)
	commentCore := core.NewCommentCore(commentRepo)

	storyHandler := handler.NewGraphQLHandler(userCore, storyCore, commentCore)
	router := mux.NewRouter()
	router.HandleFunc("/graphql", storyHandler.ServeHTTP).Methods(http.MethodPost)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)

}

func initDatabase() *mongo.Client {
	dsn := "mongodb://root:root@localhost:27017"
	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}
