package handler

import (
	"Moonlight_/core"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
)

func defineGetProjectField(projectCore core.ProjectCore, projectType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: projectType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(string)
			return projectCore.GetProject(id)
		},
	}

}

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

// func defineCreateProjectField(projectCore core.ProjectCore, projectType *graphql.Object) *graphql.Field {
// 	return &graphql.Field{
// 		Type: projectType,
// 		Args: graphql.FieldConfigArgument{
// 			"input": &graphql.ArgumentConfig{
// 				Type: graphql.NewNonNull(graphql.NewInputObject(
// 					graphql.InputObjectConfig{
// 						Name: "CreateProjectInput",
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

// 			return projectCore.CreateProject(title, content, authorID)
// 		},
// 	}
// }

func defineCreateProjectField(projectCore core.ProjectCore, projectType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: projectType,
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
			"tasks": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "TaskInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"title": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"assignedTo": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
					},
				))),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			project := core.New_project_req{
				Title:       p.Args["title"].(string),
				Description: p.Args["description"].(string),
				DueDate:     p.Args["dueDate"].(time.Time),
				Priority:    p.Args["priority"].(string),
			}
			tasksArg, ok := p.Args["tasks"].([]interface{})
			if ok && len(tasksArg) > 0 {
				var tasks []core.Task
				for _, task := range tasksArg {
					if taskMap, ok := task.(map[string]interface{}); ok {
						taskInput := core.Task{
							Title:      taskMap["title"].(string),
							AssignedTo: taskMap["assignedTo"].(string),
						}
						tasks = append(tasks, taskInput)
					} else {
						return nil, fmt.Errorf("task is not an object")
					}
				}
				project.Tasks = tasks
				fmt.Println("tasks found", project.Tasks)
			} else {
				fmt.Println("tasks not found")
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
				project.Labels = labels
				fmt.Println("labels found", project.Labels)
			}
			fmt.Println("project", project)
			projectResp, err := projectCore.NewProject(project)
			if err != nil {
				return nil, err
			}
			//fmt.Println("projectResp", projectResp)
			return projectResp, nil

		},
	}
}

func defineUpdateProjectField(projectCore core.ProjectCore, projectType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: projectType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"dueDate": &graphql.ArgumentConfig{
				Type: graphql.DateTime,
			},
			"priority": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"labels": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			project := core.New_project_req{
				Title:       p.Args["title"].(string),
				Description: p.Args["description"].(string),
				DueDate:     p.Args["dueDate"].(time.Time),
				Priority:    p.Args["priority"].(string),
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
				project.Labels = labels
			}

			tasksArg, ok := p.Args["tasks"].([]interface{})
			if ok {
				var tasks []core.Task
				for _, task := range tasksArg {
					if taskMap, ok := task.(map[string]interface{}); ok {
						taskInput := core.Task{
							Title:      taskMap["title"].(string),
							AssignedTo: taskMap["assignedTo"].(string),
						}
						tasks = append(tasks, taskInput)
					} else {
						return nil, fmt.Errorf("task is not an object")
					}
				}
				project.Tasks = tasks
			}
			projectResp, err := projectCore.NewProject(project)
			if err != nil {
				return nil, err
			}
			return projectResp, nil

		},
	}
}

func defineDeleteProjectField(projectCore core.ProjectCore, projectType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: projectType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			projectResp, err := projectCore.DeleteProject(id)
			if err != nil {
				return nil, err
			}
			return projectResp, nil
		},
	}
}

func defireCreateCommentField(commentCore core.CommentCore, commentType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: commentType,
		Args: graphql.FieldConfigArgument{
			"projectID": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"content": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"author": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			comment := core.New_comment_req{
				ProjectID: p.Args["projectID"].(string),
				Content:   p.Args["content"].(string),
				Author:    p.Args["author"].(string),
			}
			commentResp, err := commentCore.NewComment(comment)
			if err != nil {
				return nil, err
			}
			return commentResp, nil
		},
	}
}
