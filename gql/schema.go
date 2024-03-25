package gql

import (
	"Moonlight_/core"

	"github.com/graphql-go/graphql"
)

type GraphQLConfig struct {
	UserCore    core.UserCore
	StoryCore   core.ProjectCore
	TaskCore    core.TaskCore
	CommentCore core.CommentCore
}

func BuildSchema(config GraphQLConfig) *graphql.Schema {
	userType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id":       &graphql.Field{Type: graphql.String},
				"username": &graphql.Field{Type: graphql.String},
				"password": &graphql.Field{Type: graphql.String},
				"email":    &graphql.Field{Type: graphql.String},
				"token":    &graphql.Field{Type: graphql.String},
				"role":     &graphql.Field{Type: graphql.String},
				// "tasks": &graphql.Field{
				// 	Type: graphql.NewList(defineTaskType),
				// 	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// 		test := params.Source.(core.Edit_task_req)
				// 		return test.Tasks, nil
				// 	},
				// },
			},
		},
	)
	commentType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Comment",
			Fields: graphql.Fields{
				"id":      &graphql.Field{Type: graphql.String},
				"content": &graphql.Field{Type: graphql.String},
				"author":  &graphql.Field{Type: graphql.String},
			},
		},
	)
	taskType := graphql.NewObject(
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
	projectType := graphql.NewObject(
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
				"tasks":       defineGetTaskByProjectField(config.TaskCore, taskType),
			},
		},
	)
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"GetUser":     defineGetUserField(config.UserCore, userType),
				"GetUsers":    defineGetUsersField(config.UserCore, userType),
				"GetProject":  defineGetProjectField(config.StoryCore, projectType),
				"GetProjects": defineGetProjectsField(config.StoryCore, projectType),
				"GetTask":     defineGetTaskField(config.TaskCore, taskType),
				"GetTasks":    defineGetTasksField(config.TaskCore, taskType),
			},
		},
	)

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"CreateUser":    defineCreateUserField(config.UserCore, userType),
			"UpdateUser":    defineUpdateUserField(config.UserCore, userType),
			"DeleteUser":    defineDeleteUserField(config.UserCore, userType),
			"Login":         defineLoginField(config.UserCore),
			"CreateProject": defineCreateProjectField(config.StoryCore, projectType),
			"UpdateProject": defineUpdateProjectField(config.StoryCore, projectType),
			"DeleteProject": defineDeleteProjectField(config.StoryCore, projectType),
			"CreateTask":    defineCreateTaskField(config.TaskCore, taskType),
			"CreateComment": defireCreateCommentField(config.CommentCore, commentType),
			// Add other mutation fields here
		},
	},
	)

	schemaConfig := graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
		Types:    []graphql.Type{userType, projectType, taskType, commentType},
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(err)
	}
	return &schema
}
