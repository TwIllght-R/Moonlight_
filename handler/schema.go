package handler

import (
	"Moonlight_/core"

	"github.com/graphql-go/graphql"
)

type GraphQLHandler struct {
	Schema graphql.Schema
}

func buildSchema(userCore core.UserCore, storyCore core.ProjectCore, commentCore core.CommentCore) (graphql.Schema, error) {
	userType := defineUserType()
	storyType := defineProjectType()
	commentType := defineCommentType()
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"GetUser":     defineGetUserField(userCore, userType),
				"GetUsers":    defineGetUsersField(userCore, userType),
				"GetProject":  defineGetProjectField(storyCore, storyType),
				"GetProjects": defineGetProjectsField(storyCore, storyType),
			},
		},
	)

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"CreateUser":    defineCreateUserField(userCore, userType),
			"UpdateUser":    defineUpdateUserField(userCore, userType),
			"DeleteUser":    defineDeleteUserField(userCore, userType),
			"Login":         defineLoginField(userCore),
			"CreateProject": defineCreateProjectField(storyCore, storyType),
			"UpdateProject": defineUpdateProjectField(storyCore, storyType),
			"DeleteProject": defineDeleteProjectField(storyCore, storyType),
			"CreateComment": defireCreateCommentField(commentCore, commentType),
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

var defineTaskType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Task",
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.String},
			"title":      &graphql.Field{Type: graphql.String},
			"status":     &graphql.Field{Type: graphql.String},
			"assignedTo": &graphql.Field{Type: graphql.String},
		},
	},
)

func defineProjectType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Project",
			Fields: graphql.Fields{
				"id":          &graphql.Field{Type: graphql.String},
				"title":       &graphql.Field{Type: graphql.String},
				"description": &graphql.Field{Type: graphql.String},
				"dueDate":     &graphql.Field{Type: graphql.DateTime},
				"priority":    &graphql.Field{Type: graphql.String},
				"status":      &graphql.Field{Type: graphql.String},
				"labels":      &graphql.Field{Type: graphql.NewList(graphql.String)},
				"tasks":       &graphql.Field{Type: graphql.NewList(defineTaskType)},
			},
		},
	)
}

func defineUserType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id":       &graphql.Field{Type: graphql.String},
				"username": &graphql.Field{Type: graphql.String},
				"password": &graphql.Field{Type: graphql.String},
				"email":    &graphql.Field{Type: graphql.String},
				"token":    &graphql.Field{Type: graphql.String},
				"role":     &graphql.Field{Type: graphql.String},
			},
		},
	)
}

func defineCommentType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Comment",
			Fields: graphql.Fields{
				"id":      &graphql.Field{Type: graphql.String},
				"content": &graphql.Field{Type: graphql.String},
				"author":  &graphql.Field{Type: graphql.String},
			},
		},
	)
}
