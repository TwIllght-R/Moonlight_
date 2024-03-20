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
		Status:      "To Do",
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
		ID:          story.ID,
		Title:       story.Title,
		Description: story.Description,
		DueDate:     story.DueDate,
		Priority:    story.Priority,
		Status:      story.Status,
		Labels:      story.Labels,
		AssignedTo:  story.AssignedTo,
	}

	return &resp, nil
}

func (r storyCore) ListStory() (*[]Get_story_resp, error) {
	stories, err := r.storyRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var resp []Get_story_resp
	for _, story := range *stories {
		resp = append(resp, Get_story_resp{
			ID:          story.ID,
			Title:       story.Title,
			Description: story.Description,
			DueDate:     story.DueDate,
			Priority:    story.Priority,
			Status:      story.Status,
			Labels:      story.Labels,
			AssignedTo:  story.AssignedTo,
		})
	}

	return &resp, nil
}

func (r storyCore) UpdateStory(id string, req Update_story_req) (*New_story_resp, error) {
	s := repo.Story{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		Priority:    req.Priority,
		Status:      req.Status,
		Labels:      req.Labels,
		AssignedTo:  req.AssignedTo,
	}
	NewStory, err := r.storyRepo.UpdateStory(s)
	if err != nil {
		log.Panic(err)
		return nil, err

	}
	resp := New_story_resp{
		Title: NewStory.Title,
	}

	return &resp, nil
}

func (r storyCore) DeleteStory(id string) (*Get_story_resp, error) {
	story, err := r.storyRepo.DeleteStory(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp := Get_story_resp{
		ID: story.ID,
	}

	return &resp, nil
}
