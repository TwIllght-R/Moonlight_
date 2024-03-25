package core

type New_task_req struct {
	Title      string `json:"title"`
	AssignedTo string `json:"assigned_to"`
	Project_ID string `json:"project_id"`
}

type Edit_task_req struct {
	ID         string `json:"task_id"`
	Title      string `json:"title"`
	AssignedTo string `json:"assigned_to"`
}

type Task_resp struct {
	ID         string `json:"task_id"`
	Project_ID string `json:"project_id"`
	Title      string `json:"title"`
	Status     string `json:"status"`
	//AssignedTo  string `json:"assigned_to"`
	//Attachments []Attachment `json:"attachments"`
}

type TaskCore interface {
	NewTask(req New_task_req) (*Task_resp, error)
	GetTask(id string) (*Task_resp, error)
	GetTasksByProject(id string) (*[]Task_resp, error)
	GetTasks() (*[]Task_resp, error)
	EditTask(req Edit_task_req) (*Task_resp, error)
	DelTask(id string) (*Task_resp, error)
}
