package dto

type TaskRequestDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	AssigneeID  int    `json:"assignee_id"`
	ProjectID   int    `json:"project_id"`
}

type TaskResponseDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	AssigneeID  int    `json:"assignee_id"`
	ProjectID   int    `json:"project_id"`
}
