package handler

import (
	"Moonlight_/core"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

type storyHandler struct {
	storyCore core.StoryCore
	Schema    graphql.Schema
}

// func NewStoryHandler(storyCore core.StoryCore) *storyHandler {
// 	schema, err := buildStorySchema(storyCore)
// 	if err != nil {
// 		return nil
// 	}

// 	return &storyHandler{
// 		storyCore: storyCore,
// 		Schema:    schema,
// 	}
// }

// func (h *storyHandler) ServeStory(w http.ResponseWriter, r *http.Request) {
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

// 	result := executeQuery(requestBody.Query, h.Schema)
// 	json.NewEncoder(w).Encode(result)
// }

// func buildStorySchema(storyCore core.StoryCore) (graphql.Schema, error) {
// 	storyType := defineStoryType()
// 	queryType := defineQueryType(storyCore, storyType)
// 	mutationType := defineMutationType(storyCore, storyType)

// 	return graphql.NewSchema(
// 		graphql.SchemaConfig{
// 			Query:    queryType,
// 			Mutation: mutationType,
// 		},
// 	)
// }

func defineStoryType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Story",
			Fields: graphql.Fields{
				"id":          &graphql.Field{Type: graphql.String},
				"title":       &graphql.Field{Type: graphql.String},
				"description": &graphql.Field{Type: graphql.String},
				"dueDate":     &graphql.Field{Type: graphql.DateTime},
				"priority":    &graphql.Field{Type: graphql.String},
				"status":      &graphql.Field{Type: graphql.String},
				"labels":      &graphql.Field{Type: graphql.NewList(graphql.String)},
				"assignedTo":  &graphql.Field{Type: graphql.String},
			},
		},
	)
}

// func defineQueryType(storyCore core.StoryCore, storyType *graphql.Object) *graphql.Object {
// 	return graphql.NewObject(
// 		graphql.ObjectConfig{
// 			Name: "Query",
// 			Fields: graphql.Fields{
// 				"GetStory": &graphql.Field{
// 					Type:        storyType,
// 					Description: "Get story by id",
// 					Args: graphql.FieldConfigArgument{
// 						"id": &graphql.ArgumentConfig{
// 							Type: graphql.String,
// 						},
// 					},
// 					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 						id, ok := p.Args["id"].(string)
// 						if !ok {
// 							return nil, nil
// 						}
// 						story, err := storyCore.GetStory(id)
// 						if err != nil {
// 							return nil, err
// 						}
// 						return story, nil
// 					},
// 				},
// 			},
// 		},
// 	)
// }

func defineGetStoryField(storyCore core.StoryCore, storyType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: storyType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(string)
			return storyCore.GetStory(id)
		},
	}

}

// func defineMutationType(storyCore core.StoryCore, storyType *graphql.Object) *graphql.Object {
// 	return graphql.NewObject(graphql.ObjectConfig{
// 		Name: "Mutation",
// 		Fields: graphql.Fields{
// 			"CreateStory": &graphql.Field{
// 				Type:        storyType,
// 				Description: "Create new Story",
// 				Args: graphql.FieldConfigArgument{
// 					"title":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
// 					"description": &graphql.ArgumentConfig{Type: graphql.String},
// 					"dueDate":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.DateTime)},
// 					"priority":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
// 					"labels":      &graphql.ArgumentConfig{Type: graphql.NewList(graphql.String)},
// 					"assignedTo":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
// 				},
// 				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 					story := core.New_story_req{
// 						Title:       p.Args["title"].(string),
// 						Description: p.Args["description"].(string),
// 						DueDate:     p.Args["dueDate"].(time.Time),
// 						Priority:    p.Args["priority"].(string),
// 						AssignedTo:  p.Args["assignedTo"].(string),
// 					}
// 					labelsArg, ok := p.Args["labels"].([]interface{})
// 					if ok {
// 						var labels []string
// 						for _, label := range labelsArg {
// 							if str, ok := label.(string); ok {
// 								labels = append(labels, str)
// 							} else {
// 								return nil, fmt.Errorf("label is not a string")
// 							}
// 						}
// 						story.Labels = labels
// 					}

// 					storyResp, err := storyCore.NewStory(story)
// 					if err != nil {
// 						return nil, err
// 					}
// 					return storyResp, nil
// 				},
// 			},
// 		},
// 	})
// }

// func defineCreateStoryField(storyCore core.StoryCore, storyType *graphql.Object) *graphql.Field {
// 	return &graphql.Field{
// 		Type: storyType,
// 		Args: graphql.FieldConfigArgument{
// 			"input": &graphql.ArgumentConfig{
// 				Type: graphql.NewNonNull(graphql.NewInputObject(
// 					graphql.InputObjectConfig{
// 						Name: "CreateStoryInput",
// 						Fields: graphql.InputObjectConfigFieldMap{
// 							"title": &graphql.InputObjectFieldConfig{
// 								Type: graphql.NewNonNull(graphql.String),
// 							},
// 							"content": &graphql.InputObjectFieldConfig{
// 								Type: graphql.NewNonNull(graphql.String),
// 							},
// 							"authorID": &graphql.InputObjectFieldConfig{
// 								Type: graphql.NewNonNull(graphql.Int),
// 							},
// 						},
// 					},
// 				)),
// 			},
// 		},
// 		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 			input, _ := params.Args["input"].(map[string]interface{})
// 			title, _ := input["title"].(string)
// 			content, _ := input["content"].(string)
// 			authorID, _ := input["authorID"].(int)

// 			return storyCore.CreateStory(title, content, authorID)
// 		},
// 	}
// }

func defineCreateStoryField(storyCore core.StoryCore, storyType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: storyType,
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"dueDate": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"priority": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"labels": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
			"assignedTo": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			story := core.New_story_req{
				Title:       p.Args["title"].(string),
				Description: p.Args["description"].(string),
				DueDate:     p.Args["dueDate"].(time.Time),
				Priority:    p.Args["priority"].(string),
				AssignedTo:  p.Args["assignedTo"].(string),
			}
			labelsArg, ok := p.Args["labels"].([]interface{})
			if ok {
				var labels []string
				for _, label := range labelsArg {
					if str, ok := label.(string); ok {
						labels = append(labels, str)
					} else {
						return nil, fmt.Errorf("label is not a string")
					}
				}
				story.Labels = labels
			}

			storyResp, err := storyCore.NewStory(story)
			if err != nil {
				return nil, err
			}
			return storyResp, nil

		},
	}
}
func buildSchema(userCore core.UserCore, storyCore core.StoryCore) (graphql.Schema, error) {
	//userType := defineUserType()
	storyType := defineStoryType()

	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				// "GetUser": defineGetUserField(userCore, userType),
				// "GetUsers": defineGetUsersField(userCore, userType),
				// "GetStory": defineGetStoryField(storyCore, storyType),
				"GetStory": defineGetStoryField(storyCore, storyType),
			},
		},
	)

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			// "CreateUser":  defineCreateUserField(userCore, userType),
			// "Login":       defineLoginField(userCore),
			"CreateStory": defineCreateStoryField(storyCore, storyType),
			// Add other mutation fields here
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

type GraphQLHandler struct {
	Schema graphql.Schema
}

func NewGraphQLHandler(userCore core.UserCore, storyCore core.StoryCore) *GraphQLHandler {
	schema, err := buildSchema(userCore, storyCore)
	if err != nil {
		panic(err) // Handle schema initialization error
	}

	return &GraphQLHandler{
		Schema: schema,
	}
}

func (h *GraphQLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		Query string `json:"query"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Execute GraphQL query
	result := executeQuery(requestBody.Query, h.Schema)
	json.NewEncoder(w).Encode(result)
}
