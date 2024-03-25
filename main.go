package main

import (
	"Moonlight_/core"
	"Moonlight_/gql"
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
	// Set a timeout for database operations (10 seconds)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	client := initDatabase(ctx).Database("TaskSystems")
	initTimeZone()

	// graphqlHandler := handler.New(&handler.Config{
	//     Schema:   &schema,
	//     Pretty:   true,
	//     GraphiQL: true, // Enable GraphiQL for testing
	// })
	taskCollection := client.Collection("tasks")         //tasks
	userCollection := client.Collection("users")         //users
	projectCollection := client.Collection("projecties") //projects
	commentCollection := client.Collection("comments")   //comments
	//commentCollection := client.Database("TaskSystems").Collection("comments")   //comments

	//user
	userRepo := repo.NewUserRepo(userCollection)
	userCore := core.NewUserCore(userRepo)
	//task
	taskRepo := repo.NewTaskRepo(taskCollection)
	taskCore := core.NewTaskCore(taskRepo)
	//project
	projectRepo := repo.NewProjectRepo(projectCollection)
	projectCore := core.NewProjectCore(projectRepo) //project need user ,task ,comment
	//comment
	commentRepo := repo.NewCommentRepo(commentCollection)
	commentCore := core.NewCommentCore(commentRepo)

	// userSchema, _ := graphQL.NewUserSchema(userCore)
	// taskSchema, _ := graphQL.NewTaskSchema(taskCore)

	// mergedSchema, err := graphql.MergeSchemas(graphql.MergeConfig{
	// 	Schemas: []*graphql.Schema{userSchema, taskSchema},
	// })
	// if err != nil {
	// 	// handle error
	// }
	//graphqlHandler := graphQL.NewGraphQLHandler(mergedSchema)
	//	userSchema, _ := handler.NewUserSchema(userCore)
	//userHandler := handler.NewUserHandler(userSchema)
	gqlConfig := gql.GraphQLConfig{
		UserCore:    userCore,
		StoryCore:   projectCore,
		TaskCore:    taskCore,
		CommentCore: commentCore,
	}
	gqlSchema := gql.BuildSchema(gqlConfig)
	gqlHandler := handler.NewGraphQLHandler(gqlSchema)
	router := mux.NewRouter()
	router.HandleFunc("/graphql", gqlHandler.ServeHTTP).Methods(http.MethodPost)
	//router.HandleFunc("/graphql2", graphqlHandler.ServeHTTP).Methods(http.MethodPost)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)

}

func initDatabase(ctx context.Context) *mongo.Client {
	dsn := "mongodb://root:root@localhost:27017"
	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
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
