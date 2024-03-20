package handler

import (
	"Moonlight_/core"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
)

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

func defineUpdateStoryField(storyCore core.StoryCore, storyType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: storyType,
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

			"assignedTo": &graphql.ArgumentConfig{
				Type: graphql.String,
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

func defineDeleteStoryField(storyCore core.StoryCore, storyType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: storyType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			storyResp, err := storyCore.DeleteStory(id)
			if err != nil {
				return nil, err
			}
			return storyResp, nil
		},
	}
}

func defireCreateCommentField(commentCore core.CommentCore, commentType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: commentType,
		Args: graphql.FieldConfigArgument{
			"storyID": &graphql.ArgumentConfig{
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
				StoryID: p.Args["storyID"].(string),
				Content: p.Args["content"].(string),
				Author:  p.Args["author"].(string),
			}
			commentResp, err := commentCore.NewComment(comment)
			if err != nil {
				return nil, err
			}
			return commentResp, nil
		},
	}
}
