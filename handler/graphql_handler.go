package handler

import (
	"Moonlight_/core"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

func NewGraphQLHandler(userCore core.UserCore, storyCore core.ProjectCore, commentCore core.CommentCore) *GraphQLHandler {
	schema, err := buildSchema(userCore, storyCore, commentCore)
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
	result, err := executeQuery(requestBody.Query, h.Schema)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func executeQuery(query string, schema graphql.Schema) (*graphql.Result, error) {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("failed to execute query: %v", result.Errors)
	}
	return result, nil
}
