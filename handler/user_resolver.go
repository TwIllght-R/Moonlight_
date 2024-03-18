package handler

import (
	"Moonlight_/core"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

type userHandler struct {
	userCore core.UserCore
}

func NewuserHandler(userCore core.UserCore) *userHandler {
	return &userHandler{userCore: userCore}
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var schema graphql.Schema
	schema, err := h.buildSchema()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var requestBody struct {
		Query string `json:"query"`
	}
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	result := executeQuery(requestBody.Query, schema)
	json.NewEncoder(w).Encode(result)
}

func (h *userHandler) buildSchema() (graphql.Schema, error) {

	userType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"username": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"GetUser": &graphql.Field{
					Type:        userType,
					Description: "Get user by id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: h.resolveUser,
				},
				"GetUsers": &graphql.Field{
					Type:        graphql.NewList(userType), // Change userType to List type
					Description: "Get User list",
					Resolve:     h.resolveUsers,
				},
			},
		},
	)

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"Create": &graphql.Field{
				Type:        userType,
				Description: "Create new User",
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: h.resolveCreateUser,
			},
		},
	},
	)

	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		},
	)
}

func (h *userHandler) resolveUser(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if !ok {
		return nil, nil // Return nil if ID is not provided
	}

	// Retrieve user data using core logic
	user, err := h.userCore.GetUser(id)
	if err != nil {
		return nil, err // Return error if user retrieval fails
	}
	return user, nil
}

func (h *userHandler) resolveUsers(p graphql.ResolveParams) (interface{}, error) {
	// Retrieve users data using core logic
	user, err := h.userCore.GetUsers()
	if err != nil {
		return nil, err // Return error if user retrieval fails
	}
	//fmt.Println(user)
	return user, nil
}

func (h *userHandler) resolveCreateUser(p graphql.ResolveParams) (interface{}, error) {
	username, _ := p.Args["username"].(string)
	email, _ := p.Args["email"].(string)
	password, _ := p.Args["password"].(string)

	user := core.New_user_req{
		Username: username,
		Email:    email,
		Password: password,
	}

	userResp, err := h.userCore.NewUser(user)
	if err != nil {
		return nil, err // Return error if user retrieval fails
	}
	//fmt.Println(user)
	return userResp, nil
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	return result
}
