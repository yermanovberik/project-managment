package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"project-managment/internal/app/handlers"
	"project-managment/internal/app/repository"
	"project-managment/internal/app/service"
)

func SetupRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	projectRepo := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepo)
	projectHandler := handlers.NewProjectHandler(projectService)

	userRepo := repository.NewUserRepository(db)
	userService := repository.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users", handler.ListUsers).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", handler.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", handler.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id:[0-9]+}/tasks", handler.GetTasksByUser).Methods("GET")
	router.HandleFunc("/users/search", handler.GetUserByName).Queries("name", "{name}").Methods("GET")
	router.HandleFunc("/users/search", handler.GetUserByEmail).Queries("email", "{email}").Methods("GET")

	router.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.GetTaskById).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTaskById).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTaskById).Methods("DELETE")

	router.HandleFunc("/projects", projectHandler.GetAllProjects).Methods("GET")
	router.HandleFunc("/projects", projectHandler.CreateProject).Methods("POST")
	router.HandleFunc("/projects/{id}", projectHandler.GetProjectById).Methods("GET")
	router.HandleFunc("/projects/{id}", projectHandler.UpdateProjectById).Methods("PUT")
	router.HandleFunc("/projects/{id}", projectHandler.DeleteProjectById).Methods("DELETE")
	router.HandleFunc("/projects/{id}/tasks", projectHandler.GetTasksByProject).Methods("GET")
	router.HandleFunc("/projects/search", projectHandler.SearchProjects).Methods("GET")
	router.HandleFunc("/projects/search", projectHandler.GetProjectsByManager).Queries("manager", "{manager}").Methods("GET")
	return router
}
