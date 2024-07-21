package dto

type ProjectRequestDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ManagerID   int    `json:"manager_id"`
}

type ProjectResponseDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ManagerID   int    `json:"manager_id"`
}
