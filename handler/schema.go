package handler

import (
	"Moonlight_/core"

	"github.com/graphql-go/graphql"
)

type GraphQLHandler struct {
	Schema graphql.Schema
}

func buildSchema(userCore core.UserCore, storyCore core.StoryCore, commentCore core.CommentCore) (graphql.Schema, error) {
	userType := defineUserType()
	storyType := defineStoryType()
	commentType := defineCommentType()
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"GetUser": defineGetUserField(userCore, userType),
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
			"CreateStory":   defineCreateStoryField(storyCore, storyType),
			"UpdateStory":   defineUpdateStoryField(storyCore, storyType),
			"DeleteStory":   defineDeleteStoryField(storyCore, storyType),
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

func defineUserType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id":       &graphql.Field{Type: graphql.String},
				"username": &graphql.Field{Type: graphql.String},
				"email":    &graphql.Field{Type: graphql.String},
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
