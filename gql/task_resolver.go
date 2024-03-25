package gql

import (
	"Moonlight_/core"
	"fmt"

	"github.com/graphql-go/graphql"
)

func defineGetTaskField(taskCore core.TaskCore, taskType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: taskType,
		Args: graphql.FieldConfigArgument{
			"task_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["task_id"].(string)
			return taskCore.GetTask(id)
		},
	}

}

func defineGetTasksField(taskCore core.TaskCore, taskType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(taskType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return taskCore.GetTasks()
		},
	}
}

func defineCreateTaskField(taskCore core.TaskCore, taskType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: taskType,
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"assignedTo": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"project_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			title := p.Args["title"].(string)
			assignedTo := p.Args["assignedTo"].(string)
			project_id := p.Args["project_id"].(string)
			task := core.New_task_req{
				Title:      title,
				AssignedTo: assignedTo,
				Project_ID: project_id,
			}
			taskResp, err := taskCore.NewTask(task)
			if err != nil {
				return nil, err
			}
			return taskResp, nil
		},
	}
}

func defineGetTaskByProjectField(taskCore core.TaskCore, taskType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(taskType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// Try to assert p.Source to the expected type *core.Task_resp
			project, ok := p.Source.(*core.Get_project_resp)
			if !ok {
				// Handle the case where p.Source is not of the expected type
				return nil, fmt.Errorf("unexpected type %T", p.Source)
			}

			// Now that p.Source is of type *core.Task_resp, you can access its fields
			return taskCore.GetTasksByProject(project.ID)
		},
	}
}
