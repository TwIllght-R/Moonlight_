package main

import (
	"Moonlight_/core"
	"Moonlight_/handler"
	"Moonlight_/repo"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//http.Handle("/graphql", graphql.NewGraphQLServer(userResolver))
	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	// clientOptions.SetAuth(options.Credential{
	// 	Username: "root",
	// 	Password: "root",
	// })
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())
	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	fmt.Println("Connected to MongoDB!")

	// userTest := core.Login_req{
	// 	//Username: "test",
	// 	Email:    "test",
	// 	Password: "test",
	// }
	// Retrieve a user from the database
	userCollection := client.Database("lala").Collection("users")
	storyCollection := client.Database("lala").Collection("storyies")

	userRepo := repo.NewUserRepo(userCollection)
	userCore := core.NewUserCore(userRepo)
	userHandler := handler.NewUserHandler(userCore)
	_ = userHandler
	storyRepo := repo.NewStoryRepo(storyCollection)
	storyCore := core.NewStoryCore(storyRepo)
	storyHandler := handler.NewGraphQLHandler(userCore, storyCore)
	router := mux.NewRouter()

	//http.HandleFunc("/user", userHandler.User)
	router.HandleFunc("/graphql", storyHandler.ServeHTTP).Methods(http.MethodPost)
	//router.HandleFunc("/graphql2", storyHandler.ServeStory).Methods(http.MethodPost)
	// user, err := userCore.NewUser(userTest)
	// if err != nil {
	// 	log.Print(err)
	// }
	// _ = user
	// login, err := userCore.LoginUser(userTest)
	// if err != nil {
	// 	log.Panicln(err)
	// }

	// println(*login)
	// storytest := repo.Story{
	// 	Title: "My Awesome Story",
	// 	Content: map[int]struct {
	// 		Text string `bson:"text"`
	// 		Type string `bson:"type"`
	// 	}{
	// 		1: {Text: "This is the first section with text content.", Type: "text"},
	// 		2: {Text: "Here's an image!", Type: "image"}, // You'll need to store the image URL or reference here
	// 	},
	// }
	// _ = storytest

	//story, err := storyRepo.CreateStory(storytest)
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// objectID, err := primitive.ObjectIDFromHex("65f6c332877853017abf10cc")
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// storyGet, err := storyRepo.GetStoryByID(objectID)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(storyGet.Content[1].Text)
	//_ = story
	// args := struct{ ID string }{ID: "loloo"}
	// resolver := handler.NewResolver(userCore)
	// test, err := resolver.User(args)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(test)
	// Load GraphQL schema
	// http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
	// 	result := executeQuery(r.URL.Query().Get("query"), schema)
	// 	json.NewEncoder(w).Encode(result)
	// })

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)

}
