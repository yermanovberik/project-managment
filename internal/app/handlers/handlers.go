package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"project-managment/internal/app/dto"
	"project-managment/internal/app/models"
	"project-managment/internal/app/repository"
	"strconv"
)

type Handler struct {
	Repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.Repo.CreateUser(user)

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Repo.GetAllUsers()

	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	var userResponse []dto.UserResponseDTO

	for _, u := range users {
		userResponse = append(userResponse, dto.UserResponseDTO{
			ID:               int(u.ID),
			Name:             u.Name,
			Email:            u.Email,
			RegistrationDate: u.RegistrationDate,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(userResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		return
	}

	user, err := h.Repo.GetById(id)

	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}
	var userResponse dto.UserResponseDTO

	userResponse = dto.UserResponseDTO{
		ID:               int(user.ID),
		Name:             user.Name,
		Email:            user.Email,
		RegistrationDate: user.RegistrationDate,
	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(userResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.Repo.DeleteUserById(id)

	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted succesfully"))
}

func (h *Handler) GetUserByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	user, err := h.Repo.GetUserByName(name)

	if err != nil {
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

func (h *Handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	user, err := h.Repo.GetUserByEmail(email)

	if err != nil {
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

func (h *Handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task dto.TaskRequestDTO
	err := json.NewDecoder(r.Body).Decode(task)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.Repo.GetById(task.AssigneeID)

}

func (h *Handler) GetTasksByUser(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) GetTaskById(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) UpdateTaskById(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) DeleteTaskById(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	project, err := h.Repo.GetAllProjects()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(project); err != nil {
		http.Error(w, "Failed to get projects", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var project dto.ProjectRequestDTO
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		http.Error(w, "Invalid request for project", http.StatusBadRequest)
		return
	}
	err = h.Repo.CreateProject(project)

	if err != nil {
		http.Error(w, "Failed to create project", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetProjectById(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) UpdateProjectById(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) DeleteProjectById(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) GetTasksByProject(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) SearchProjects(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) GetProjectsByManager(writer http.ResponseWriter, request *http.Request) {

}
