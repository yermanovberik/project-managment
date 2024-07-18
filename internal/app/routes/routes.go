package routes

import (
	"github.com/gorilla/mux"
	"project-managment/internal/app/handlers"
)

func SetupRoutes(handler *handlers.Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", handler.CreateUser).Methods("POST")
	router.HandleFunc("/users", handler.ListUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handler.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}/tasks", handler.GetTasksByUser).Methods("GET")
	router.HandleFunc("/users/search/{name}", handler.GetUserByName).Methods("GET")
	router.HandleFunc("/users/search/{email}", handler.GetUserByEmail).Methods("GET")

	router.HandleFunc("/tasks", handler.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", handler.GetTaskById).Methods("GET")
	router.HandleFunc("/tasks/{id}", handler.UpdateTaskById).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handler.DeleteTaskById).Methods("DELETE")

	router.HandleFunc("/projects", handler.GetAllProjects).Methods("GET")
	router.HandleFunc("/projects", handler.CreateProject).Methods("POST")
	router.HandleFunc("/projects/{id}", handler.GetProjectById).Methods("GET")
	router.HandleFunc("/projects/{id}", handler.UpdateProjectById).Methods("PUT")
	router.HandleFunc("/projects/{id}", handler.DeleteProjectById).Methods("DELETE")
	router.HandleFunc("/projects/{id}/tasks", handler.GetTasksByProject).Methods("GET")
	router.HandleFunc("/projects/search", handler.SearchProjects).Methods("GET")
	router.HandleFunc("/projects/search", handler.GetProjectsByManager).Queries("manager", "{manager}").Methods("GET")
	return router
}
