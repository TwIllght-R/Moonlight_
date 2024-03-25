package core

import (
	"Moonlight_/repo"
	"log"
	"time"
)

type projectCore struct {
	projectRepo repo.ProjectRepo
}

func NewProjectCore(projectRepo repo.ProjectRepo) ProjectCore {
	return projectCore{projectRepo: projectRepo}
}

func (r projectCore) NewProject(req New_project_req) (*New_project_resp, error) {
	// Validate Priority
	if errVal := ValidatePriority(req.Priority); errVal != nil {
		return nil, errVal
	}
	// if errVal := ValidateStatus(req.Status); errVal != nil {
	// 	return nil, errVal
	// }
	// if errVal := ValidateLabels(req.Labels); errVal != nil {
	// 	return nil, errVal
	// }
	// if errVal := ValidateDate(req.DueDate); errVal != nil {
	// 	return nil, errVal
	// }
	//check if user exists

	// var attachments []repo.Attachment
	// for _, attachment := range req.Attachments {
	// 	attachments = append(attachments, repo.Attachment(attachment))
	// }

	// var tasks []repo.Task
	// for _, v := range req.Tasks {
	//     tasks = append(tasks, repo.Task{
	//         Title:      v.Title,
	//         AssignedTo: v.AssignedTo,
	//         Status:     "To Do",
	//     })
	// }

	// preallocate the slice for performance
	// tasks := make([]repo.Task, len(req.Tasks))
	// for i, v := range req.Tasks {
	// 	tasks[i] = repo.Task{
	// 		Title:      v.Title,
	// 		AssignedTo: v.AssignedTo,
	// 		Status:     "To Do",
	// 	}
	// }

	//check err

	s := repo.Project{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		Priority:    req.Priority,
		Labels:      req.Labels,
		Status:      "To Do",
		//AssignedTo:  req.AssignedTo,
		//Attachments: attachments,
		//Tasks: tasks,
	}
	NewProject, err := r.projectRepo.CreateProject(s)
	if err != nil {
		log.Panic(err)
		return nil, err

	}
	// newTasks := make([]Task, len(NewProject.Tasks))
	// for i, v := range NewProject.Tasks {
	// 	newTasks[i] = Task{
	// 		Title:      v.Title,
	// 		Status:     v.Status,
	// 		AssignedTo: v.AssignedTo,
	// 	}
	// }
	resp := New_project_resp{
		Title:       NewProject.Title,
		Description: NewProject.Description,
		DueDate:     NewProject.DueDate,
		Priority:    NewProject.Priority,
		Labels:      NewProject.Labels,
		//	AssignedTo:  NewProject.AssignedTo,
		//Attachments: attachments,
		//Tasks: newTasks, // Use the converted tasks slice
	}

	return &resp, nil

}

func (r projectCore) GetProject(id string) (*Get_project_resp, error) {
	project, err := r.projectRepo.GetProjectByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var attachments []Attachment
	for _, attachment := range project.Attachments {
		attachments = append(attachments, Attachment(attachment))
	}
	// newTasks := make([]Task, len(project.Tasks))
	// for i, v := range project.Tasks {
	// 	newTasks[i] = Task{
	// 		Title:      v.Title,
	// 		AssignedTo: v.AssignedTo,
	// 	}
	// }
	resp := Get_project_resp{
		ID:          ConvertObjectIDToString(project.ID),
		Title:       project.Title,
		Description: project.Description,
		DueDate:     project.DueDate,
		Priority:    project.Priority,
		Status:      project.Status,
		Labels:      project.Labels,
		//AssignedTo:  project.AssignedTo,
		Attachments: attachments,
		//Tasks:       newTasks,
	}

	return &resp, nil
}

func (r projectCore) ListProject() (*[]Get_project_resp, error) {
	stories, err := r.projectRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var resp []Get_project_resp
	for _, project := range *stories {
		resp = append(resp, Get_project_resp{
			ID:          ConvertObjectIDToString(project.ID),
			Title:       project.Title,
			Description: project.Description,
			DueDate:     project.DueDate,
			Priority:    project.Priority,
			Status:      project.Status,
			Labels:      project.Labels,
			//	AssignedTo:  project.AssignedTo,
		})
	}

	return &resp, nil
}

func (r projectCore) UpdateProject(req Update_project_req) (*New_project_resp, error) {
	pid, err := ConvertStringToObjectID(req.ID)
	if err != nil {
		return nil, err
	}
	s := repo.Project{
		ID:          pid,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		Priority:    req.Priority,
		Status:      req.Status,
		Labels:      req.Labels,
		UpdatedAt:   time.Now(),
		//	AssignedTo:  req.AssignedTo,
	}
	NewProject, err := r.projectRepo.UpdateProject(s)
	if err != nil {
		log.Panic(err)
		return nil, err

	}
	resp := New_project_resp{
		Title: NewProject.Title,
	}

	return &resp, nil
}

func (r projectCore) DeleteProject(id string) (*Get_project_resp, error) {
	project, err := r.projectRepo.DeleteProject(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp := Get_project_resp{
		ID: ConvertObjectIDToString(project.ID),
	}

	return &resp, nil
}
