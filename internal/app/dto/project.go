package dto

type ProjectRequestDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ManagerID   int    `json:"manager_id"`
}
