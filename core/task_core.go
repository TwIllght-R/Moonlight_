package core

import (
	"Moonlight_/repo"
	"log"
	"time"
)

type taskCore struct {
	taskRepo repo.TaskRepo
}

func NewTaskCore(taskRepo repo.TaskRepo) TaskCore {
	return taskCore{taskRepo: taskRepo}
}

func (r taskCore) NewTask(req New_task_req) (*Task_resp, error) {

	id, err := ConvertStringToObjectID(req.Project_ID)
	if err != nil {
		return nil, err
	}
	t := repo.Task{
		Title:      req.Title,
		AssignedTo: req.AssignedTo,
		Status:     "todo",
		CreatedAt:  time.Now(),
		Project_ID: id,
	}
	newTasK, err := r.taskRepo.CreateTask(t)
	if err != nil {
		return nil, err
	}
	resp := Task_resp{
		// ID:         newTasK.ID,
		// Project_ID: newTasK.Project_ID,
		Title:  newTasK.Title,
		Status: newTasK.Status,
	}
	return &resp, nil
}

func (r taskCore) GetTask(id string) (*Task_resp, error) {
	return nil, nil
}

func (r taskCore) GetTasksByProject(project_id string) (*[]Task_resp, error) {
	pid, err := ConvertStringToObjectID(project_id)
	if err != nil {
		return nil, err
	}
	tasks, err := r.taskRepo.GetTaskByProjectID(pid)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var resp []Task_resp
	for _, task := range *tasks {
		resp = append(resp, Task_resp{
			// ID:         task.ID,
			// Project_ID: task.Project_ID,
			Title:  task.Title,
			Status: task.Status,
		})
	}
	return &resp, nil

}

func (r taskCore) GetTasks() (*[]Task_resp, error) {
	return nil, nil
}

func (r taskCore) EditTask(req Edit_task_req) (*Task_resp, error) {
	return nil, nil
}

func (r taskCore) DelTask(id string) (*Task_resp, error) {
	return nil, nil
}
