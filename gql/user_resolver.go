package gql

import (
	"Moonlight_/core"

	"github.com/graphql-go/graphql"
)

// type userHandler struct {
// 	userCore core.UserCore
// 	Schema   graphql.Schema
// }

// func NewUserHandler(userCore core.UserCore) *userHandler {
// 	// Define GraphQL schema
// 	schema, err := buildUserSchema(userCore)
// 	if err != nil {
// 		panic(err) // Handle schema initialization error
// 	}

// 	return &userHandler{
// 		userCore: userCore,
// 		Schema:   schema,
// 	}
// }

// func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var requestBody struct {
// 		Query string `json:"query"`
// 	}
// 	err := json.NewDecoder(r.Body).Decode(&requestBody)
// 	if err != nil {
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	// Execute GraphQL query
// 	result := executeQuery(requestBody.Query, h.Schema)
// 	json.NewEncoder(w).Encode(result)
// }

// func buildUserSchema(userCore core.UserCore) (graphql.Schema, error) {
// 	userType := graphql.NewObject(
// 		graphql.ObjectConfig{
// 			Name: "User",
// 			Fields: graphql.Fields{
// 				"id": &graphql.Field{
// 					Type: graphql.String,
// 				},
// 				"username": &graphql.Field{
// 					Type: graphql.String,
// 				},
// 				"email": &graphql.Field{
// 					Type: graphql.String,
// 				},
// 				"password": &graphql.Field{
// 					Type: graphql.String,
// 				},
// 			},
// 		},
// 	)

// 	queryType := graphql.NewObject(
// 		graphql.ObjectConfig{
// 			Name: "Query",
// 			Fields: graphql.Fields{
// 				"GetUser": &graphql.Field{
// 					Type:        userType,
// 					Description: "Get user by id",
// 					Args: graphql.FieldConfigArgument{
// 						"id": &graphql.ArgumentConfig{
// 							Type: graphql.String,
// 						},
// 					},
// 					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 						id, ok := p.Args["id"].(string)
// 						if !ok {
// 							return nil, nil // Return nil if ID is not provided
// 						}
// 						// Retrieve user data using core logic
// 						user, err := userCore.GetUser(id)
// 						if err != nil {
// 							return nil, err // Return error if user retrieval fails
// 						}
// 						return user, nil
// 					},
// 				},
// 				"GetUsers": &graphql.Field{
// 					Type:        graphql.NewList(userType),
// 					Description: "Get User list",
// 					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 						// Retrieve users data using core logic
// 						users, err := userCore.GetUsers()
// 						if err != nil {
// 							return nil, err // Return error if user retrieval fails
// 						}
// 						return users, nil
// 					},
// 				},
// 			},
// 		},
// 	)

// 	mutationType := graphql.NewObject(graphql.ObjectConfig{
// 		Name: "Mutation",
// 		Fields: graphql.Fields{
// 			"CreateUser": &graphql.Field{
// 				Type:        userType,
// 				Description: "Create new User",
// 				Args: graphql.FieldConfigArgument{
// 					"username": &graphql.ArgumentConfig{
// 						Type: graphql.NewNonNull(graphql.String),
// 					},
// 					"email": &graphql.ArgumentConfig{
// 						Type: graphql.String,
// 					},
// 					"password": &graphql.ArgumentConfig{
// 						Type: graphql.NewNonNull(graphql.String),
// 					},
// 				},
// 				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 					username, _ := p.Args["username"].(string)
// 					email, _ := p.Args["email"].(string)
// 					password, _ := p.Args["password"].(string)

// 					user := core.New_user_req{
// 						Username: username,
// 						Email:    email,
// 						Password: password,
// 					}

// 					userResp, err := userCore.NewUser(user)
// 					if err != nil {
// 						return nil, err // Return error if user creation fails
// 					}
// 					return userResp, nil
// 				},
// 			},
// 			"Login": &graphql.Field{
// 				Type:        graphql.String,
// 				Description: "Login user and get authentication token",
// 				Args: graphql.FieldConfigArgument{
// 					"email": &graphql.ArgumentConfig{
// 						Type: graphql.String,
// 					},
// 					"password": &graphql.ArgumentConfig{
// 						Type: graphql.NewNonNull(graphql.String),
// 					},
// 				},
// 				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 					email, _ := p.Args["email"].(string)
// 					password, _ := p.Args["password"].(string)

// 					user := core.Login_req{
// 						Email:    email,
// 						Password: password,
// 					}

// 					token, err := userCore.LoginUser(user)
// 					if err != nil {
// 						return nil, err // Return error if user creation fails
// 					}
// 					return token, nil
// 				},
// 			},
// 		},
// 	},
// 	)

//		return graphql.NewSchema(
//			graphql.SchemaConfig{
//				Query:    queryType,
//				Mutation: mutationType,
//			},
//		)
//	}
func defineGetUserField(userCore core.UserCore, userType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(string)
			return userCore.GetUser(id)
		},
	}

}
func defineCreateUserField(userCore core.UserCore, userType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: userType,
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
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			username := p.Args["username"].(string)
			email := p.Args["email"].(string)
			password := p.Args["password"].(string)
			user := core.New_user_req{
				Username: username,
				Email:    email,
				Password: password,
			}
			userResp, err := userCore.NewUser(user)
			if err != nil {
				return nil, err
			}
			return userResp, nil
		},
	}
}

func defineLoginField(userCore core.UserCore) *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			email := p.Args["email"].(string)
			password := p.Args["password"].(string)
			user := core.Login_req{
				Email:    email,
				Password: password,
			}
			token, err := userCore.LoginUser(user)
			if err != nil {
				return nil, err
			}
			return token, nil
		},
	}
}

func defineGetUsersField(userCore core.UserCore, userType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(userType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			users, err := userCore.GetUsers()
			if err != nil {
				return nil, err
			}
			return users, nil
		},
	}
}
func defineUpdateUserField(userCore core.UserCore, userType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"user_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"username": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			user_id := p.Args["user_id"].(string)
			username := p.Args["username"].(string)
			email := p.Args["email"].(string)
			password := p.Args["password"].(string)
			user := core.Edit_User_req{
				User_Id:  user_id,
				Username: username,
				Email:    email,
				Password: password,
			}
			userResp, err := userCore.EditUser(user)
			if err != nil {
				return nil, err
			}
			return userResp, nil
		},
	}
}
func defineDeleteUserField(userCore core.UserCore, userType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			userResp, err := userCore.DelUser(id)
			if err != nil {
				return nil, err
			}
			return userResp, nil
		},
	}
}
