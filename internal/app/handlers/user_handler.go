package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"project-managment/internal/app/dto"
	"project-managment/internal/app/models"
	"project-managment/internal/app/repository"
	"project-managment/internal/app/service"
	"strconv"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	var userResponses []dto.UserResponseDTO
	for _, u := range users {
		userResponses = append(userResponses, dto.UserResponseDTO{
			ID:               u.ID,
			Name:             u.Name,
			Email:            u.Email,
			RegistrationDate: u.RegistrationDate,
			Role:             u.Role,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userResponses); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest dto.UserRequestDTO
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	var user models.User

	user = models.User{
		Name:             userRequest.Name,
		Email:            userRequest.Email,
		RegistrationDate: userRequest.RegistrationDate,
		Role:             userRequest.Role,
	}
	err = h.Service.CreateUser(user)

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetById(id)

	if err != nil {
		if err == repository.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	userResponse := dto.UserResponseDTO{
		ID:               int(user.ID),
		Name:             user.Name,
		Email:            user.Email,
		RegistrationDate: user.RegistrationDate,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.Service.DeleteUserById(id)

	if err != nil {
		if err == repository.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusInternalServerError)
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted succesfully"))
}

func (h *UserHandler) GetUserByName(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	user, err := h.Service.GetUserByName(name)

	if err != nil {
		if err == repository.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get user", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusBadRequest)
		return
	}
}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	user, err := h.Service.GetUserByEmail(email)
	if err == repository.ErrUserNotFound {
		if err == repository.ErrUserNotFound {
			http.Error(w, "User with this email not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get user", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusBadRequest)
		return
	}
}
