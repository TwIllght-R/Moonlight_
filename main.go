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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var schema, _ = graphql.NewSchema(
// 	graphql.SchemaConfig{
// 		Query: queryType,
// 		//Mutation: mutationType,
// 	},
// )
// var products = []Product{
// 	{
// 		ID:    1,
// 		Name:  "Chicha Morada",
// 		Info:  "Chicha morada is a beverage originated in the Andean regions of Perú but is actually consumed at a national level (wiki)",
// 		Price: 7.99,
// 	},
// 	{
// 		ID:    2,
// 		Name:  "Chicha de jora",
// 		Info:  "Chicha de jora is a corn beer chicha prepared by germinating maize, extracting the malt sugars, boiling the wort, and fermenting it in large vessels (traditionally huge earthenware vats) for several days (wiki)",
// 		Price: 5.95,
// 	},
// 	{
// 		ID:    3,
// 		Name:  "Pisco",
// 		Info:  "Pisco is a colorless or yellowish-to-amber colored brandy produced in winemaking regions of Peru and Chile (wiki)",
// 		Price: 9.95,
// 	},
// }

// type Product struct {
// 	ID    int64   `json:"id"`
// 	Name  string  `json:"name"`
// 	Info  string  `json:"info,omitempty"`
// 	Price float64 `json:"price"`
// }

// var productType = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name: "Product",
// 		Fields: graphql.Fields{
// 			"id": &graphql.Field{
// 				Type: graphql.Int,
// 			},
// 			"name": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"info": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"price": &graphql.Field{
// 				Type: graphql.Float,
// 			},
// 		},
// 	},
// )
// var queryType = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name: "Query",
// 		Fields: graphql.Fields{
// 			/* Get (read) single product by id
// 			   http://localhost:8080/product?query={product(id:1){name,info,price}}
// 			*/
// 			"product": &graphql.Field{
// 				Type:        productType,
// 				Description: "Get product by id",
// 				Args: graphql.FieldConfigArgument{
// 					"id": &graphql.ArgumentConfig{
// 						Type: graphql.Int,
// 					},
// 				},
// 				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 					id, ok := p.Args["id"].(int)
// 					if ok {
// 						// Find product
// 						for _, product := range products {
// 							if int(product.ID) == id {
// 								return product, nil
// 							}
// 						}
// 					}
// 					return nil, nil
// 				},
// 			},
// 			/* Get (read) product list
// 			   http://localhost:8080/product?query={list{id,name,info,price}}
// 			*/
// 			"list": &graphql.Field{
// 				Type:        graphql.NewList(productType),
// 				Description: "Get product list",
// 				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 					return products, nil
// 				},
// 			},
// 		},
// 	})

// func executeQuery(query string, schema graphql.Schema) *graphql.Result {
// 	result := graphql.Do(graphql.Params{
// 		Schema:        schema,
// 		RequestString: query,
// 	})
// 	if len(result.Errors) > 0 {
// 		fmt.Printf("errors: %v", result.Errors)
// 	}
// 	return result
// }

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

	userTest := core.Login_req{
		//Username: "test",
		Email:    "test",
		Password: "test",
	}
	// Retrieve a user from the database
	userCollection := client.Database("lala").Collection("users")
	storyCollection := client.Database("lala").Collection("storyies")

	userRepo := repo.NewUserRepo(userCollection)
	userCore := core.NewUserCore(userRepo)
	userHandler := handler.NewuserHandler(userCore)
	storyRepo := repo.NewStoryRepo(storyCollection)
	_ = storyRepo
	router := mux.NewRouter()

	//http.HandleFunc("/user", userHandler.User)
	router.HandleFunc("/graphql", userHandler.ServeHTTP).Methods(http.MethodPost)
	// user, err := userCore.NewUser(userTest)
	// if err != nil {
	// 	log.Print(err)
	// }
	// _ = user
	login, err := userCore.LoginUser(userTest)
	if err != nil {
		log.Panicln(err)
	}

	println(*login)
	storytest := repo.Story{
		Title: "My Awesome Story",
		Content: map[int]struct {
			Text string `bson:"text"`
			Type string `bson:"type"`
		}{
			1: {Text: "This is the first section with text content.", Type: "text"},
			2: {Text: "Here's an image!", Type: "image"}, // You'll need to store the image URL or reference here
		},
	}
	_ = storytest

	//story, err := storyRepo.CreateStory(storytest)
	if err != nil {
		log.Panicln(err)
	}
	objectID, err := primitive.ObjectIDFromHex("65f6c332877853017abf10cc")
	if err != nil {
		log.Panicln(err)
	}
	storyGet, err := storyRepo.GetStoryByID(objectID)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(storyGet.Content[1].Text)
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
