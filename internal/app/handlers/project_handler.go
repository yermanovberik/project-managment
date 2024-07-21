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

type ProjectHandler struct {
	Service service.ProjectService
}

func NewProjectHandler(service service.ProjectService) *ProjectHandler {
	return &ProjectHandler{Service: service}
}

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var projectRequest dto.ProjectRequestDTO
	err := json.NewDecoder(r.Body).Decode(&projectRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	project := models.Project{
		Title:       projectRequest.Title,
		Description: projectRequest.Description,
		ManagerID:   projectRequest.ManagerID,
	}

	err = h.Service.CreateProject(project)
	if err != nil {
		http.Error(w, "Failed to create project", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(project); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *ProjectHandler) GetProjectById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		return
	}

	project, err := h.Service.GetProjectById(id)
	if err != nil {
		if err == repository.ErrProjectNotFound {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get project", http.StatusInternalServerError)
		return
	}

	projectResponse := dto.ProjectResponseDTO{
		ID:          project.ID,
		Title:       project.Title,
		Description: project.Description,
		ManagerID:   project.ManagerID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(projectResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *ProjectHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.Service.GetAllProjects()
	if err != nil {
		http.Error(w, "Failed to get projects", http.StatusInternalServerError)
		return
	}

	var projectResponses []dto.ProjectResponseDTO
	for _, p := range projects {
		projectResponses = append(projectResponses, dto.ProjectResponseDTO{
			ID:          p.ID,
			Title:       p.Title,
			Description: p.Description,
			ManagerID:   p.ManagerID,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(projectResponses); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		return
	}

	var projectRequest dto.ProjectRequestDTO
	err = json.NewDecoder(r.Body).Decode(&projectRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	project := models.Project{
		ID:          id,
		Title:       projectRequest.Title,
		Description: projectRequest.Description,
		ManagerID:   projectRequest.ManagerID,
	}

	err = h.Service.UpdateProject(project)
	if err != nil {
		if err == repository.ErrProjectNotFound {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to update project", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(project); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *ProjectHandler) DeleteProjectById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteProjectById(id)
	if err != nil {
		if err == repository.ErrProjectNotFound {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Project deleted successfully"))
}

func (h *ProjectHandler) GetTasksByProjectId(w http.ResponseWriter, r *http.Request) {
	projectId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		return
	}

	tasks, err := h.Service.GetTasksByProjectId(projectId)
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}

	var taskResponses []dto.TaskResponseDTO
	for _, t := range tasks {
		taskResponses = append(taskResponses, dto.TaskResponseDTO{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(taskResponses); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *ProjectHandler) SearchProjectsByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	projects, err := h.Service.SearchProjectsByTitle(title)
	if err != nil {
		http.Error(w, "Failed to search projects", http.StatusInternalServerError)
		return
	}

	var projectResponses []dto.ProjectResponseDTO
	for _, p := range projects {
		projectResponses = append(projectResponses, dto.ProjectResponseDTO{
			ID:          p.ID,
			Title:       p.Title,
			Description: p.Description,
			ManagerID:   p.ManagerID,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(projectResponses); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *ProjectHandler) SearchProjectsByManager(w http.ResponseWriter, r *http.Request) {
	managerId, err := strconv.Atoi(r.URL.Query().Get("manager"))
	if err != nil {
		http.Error(w, "Manager ID is not a number", http.StatusBadRequest)
		return
	}

	projects, err := h.Service.SearchProjectsByManager(managerId)
	if err != nil {
		http.Error(w, "Failed to search projects", http.StatusInternalServerError)
		return
	}

	var projectResponses []dto.ProjectResponseDTO
	for _, p := range projects {
		projectResponses = append(projectResponses, dto.ProjectResponseDTO{
			ID:          p.ID,
			Title:       p.Title,
			Description: p.Description,
			ManagerID:   p.ManagerID,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(projectResponses); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
