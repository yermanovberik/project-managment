package dto

type UserRequestDTO struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	RegistrationDate string `json:"registration_date"`
	Role             string `json:"role"`
}

type UserResponseDTO struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	RegistrationDate string `json:"registration_date"`
	Role             string `json:"role"`
}
