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

type TaskHandler struct {
	Service service.TaskService
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskRequest dto.TaskRequestDTO
	err := json.NewDecoder(r.Body).Decode(&taskRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	task := models.Task{
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
		Status:      taskRequest.Status,
		Priority:    taskRequest.Priority,
		AssigneeID:  taskRequest.AssigneeID,
		ProjectID:   taskRequest.ProjectID,
	}

	err = h.Service.CreateTask(task)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *TaskHandler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		return
	}

	task, err := h.Service.GetTaskById(id)
	if err != nil {
		if err == repository.ErrTaskNotFound {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get task", http.StatusInternalServerError)
		return
	}

	taskResponse := dto.TaskResponseDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		AssigneeID:  task.AssigneeID,
		ProjectID:   task.ProjectID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(taskResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetAllTasks()
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
			Priority:    t.Priority,
			AssigneeID:  t.AssigneeID,
			ProjectID:   t.ProjectID,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(taskResponses); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		return
	}

	var taskRequest dto.TaskRequestDTO
	err = json.NewDecoder(r.Body).Decode(&taskRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	task, err := h.Service.GetTaskById(id)
	if err != nil {
		if err == repository.ErrTaskNotFound {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get task", http.StatusInternalServerError)
		return
	}

	task.Title = taskRequest.Title
	task.Description = taskRequest.Description
	task.Status = taskRequest.Status
	task.Priority = taskRequest.Priority
	task.AssigneeID = taskRequest.AssigneeID
	task.ProjectID = taskRequest.ProjectID

	err = h.Service.UpdateTask(task)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	taskResponse := dto.TaskResponseDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		AssigneeID:  task.AssigneeID,
		ProjectID:   task.ProjectID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(taskResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *TaskHandler) DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteTaskById(id)
	if err != nil {
		if err == repository.ErrTaskNotFound {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task deleted successfully"))
}

func (h *TaskHandler) SearchTasksByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	tasks, err := h.Service.SearchTasksByTitle(title)
	if err != nil {
		http.Error(w, "Failed to search tasks", http.StatusInternalServerError)
		return
	}

	var taskResponses []dto.TaskResponseDTO
	for _, t := range tasks {
		taskResponses = append(taskResponses, dto.TaskResponseDTO{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			Priority:    t.Priority,
			AssigneeID:  t.AssigneeID,
			ProjectID:   t.ProjectID,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(taskResponses); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *TaskHandler) SearchTasksByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	tasks, err := h.Service.SearchTasksByStatus(status)
	if err != nil {
		http.Error(w, "Failed to search tasks", http.StatusInternalServerError)
		return
	}

	var taskResponses []dto.TaskResponseDTO
	for _, t := range tasks {
		taskResponses = append(taskResponses, dto.TaskResponseDTO{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			Priority:    t.Priority,
			AssigneeID:  t.AssigneeID,
			ProjectID:   t.ProjectID,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(taskResponses); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *TaskHandler) SearchTasksByPriority(w http.ResponseWriter, r *http.Request) {
	priority := r.URL.Query().Get("priority")

	tasks, err := h.Service.SearchTasksByPriority(priority)
	if err != nil {
		http.Error(w, "Failed to search tasks", http.StatusInternalServerError)
		return
	}

	var taskResponses []dto.TaskResponseDTO
	for _, t := range tasks {
		taskResponses = append(taskResponses, dto.TaskResponseDTO{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			Priority:    t.Priority,
			AssigneeID:  t.AssigneeID,
			ProjectID:   t.ProjectID,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(taskResponses); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *TaskHandler) SearchTasksByAssignee(w http.ResponseWriter, r *http.Request) {
	assigneeId, err := strconv.Atoi(r.URL.Query().Get("assignee"))
	if err != nil {
		http.Error(w, "Assignee ID is not a number", http.StatusBadRequest)
		return
	}

	tasks, err := h.Service.SearchTasksByAssignee(assigneeId)
	if err != nil {
		http.Error(w, "Failed to search tasks", http.StatusInternalServerError)
		return
	}

	var taskResponses []dto.TaskResponseDTO
	for _, t := range tasks {
		taskResponses = append(taskResponses, dto.TaskResponseDTO{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			Priority:    t.Priority,
			AssigneeID:  t.AssigneeID,
			ProjectID:   t.ProjectID,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(taskResponses); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *TaskHandler) SearchTasksByProject(w http.ResponseWriter, r *http.Request) {
	projectId, err := strconv.Atoi(r.URL.Query().Get("project"))
	if err != nil {
		http.Error(w, "Project ID is not a number", http.StatusBadRequest)
		return
	}

	tasks, err := h.Service.SearchTasksByProject(projectId)
	if err != nil {
		http.Error(w, "Failed to search tasks", http.StatusInternalServerError)
		return
	}

	var taskResponses []dto.TaskResponseDTO
	for _, t := range tasks {
		taskResponses = append(taskResponses, dto.TaskResponseDTO{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			Priority:    t.Priority,
			AssigneeID:  t.AssigneeID,
			ProjectID:   t.ProjectID,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(taskResponses); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
