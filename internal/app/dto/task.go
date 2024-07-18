package dto

type TaskRequestDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
	AssigneeID  int    `json:"assignee_id"`
	ProjectID   int    `json:"project_id"`
}
