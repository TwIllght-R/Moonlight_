package core

import (
	"Moonlight_/repo"
	"log"
)

type storyCore struct {
	storyRepo repo.StoryRepo
}

func NewStoryCore(storyRepo repo.StoryRepo) StoryCore {
	return storyCore{storyRepo: storyRepo}
}

func (r storyCore) NewStory(req New_story_req) (*New_story_resp, error) {
	s := repo.Story{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		Priority:    req.Priority,
		Labels:      req.Labels,
		AssignedTo:  req.AssignedTo,
	}
	NewStory, err := r.storyRepo.CreateStory(s)
	if err != nil {
		log.Panic(err)
		return nil, err

	}
	resp := New_story_resp{
		Title: NewStory.Title,
	}

	return &resp, nil

}

func (r storyCore) GetStory(id string) (*Get_story_resp, error) {
	story, err := r.storyRepo.GetStoryByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp := Get_story_resp{
		ID: story.ID,
	}

	return &resp, nil
}
