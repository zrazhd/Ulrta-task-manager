package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	httpHandler "github.com/zrazhd/Ulrta-task-manager/internal/delivery/http"
	repository "github.com/zrazhd/Ulrta-task-manager/internal/repository/postgres"
	"github.com/zrazhd/Ulrta-task-manager/internal/usecase"
)

func main() {
	ctx := context.Background()
	connStr := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Cannot connect to Database: %s", err)
	}

	repo := repository.NewTaskRepo(pool)
	taskService := usecase.NewTaskService(repo)
	taskHandler := httpHandler.NewTaskHandler(taskService)

	userRepo := repository.NewUserRepo(pool)
	userService := usecase.NewUserService(userRepo)
	userHandler := httpHandler.NewUserHandler(userService)

	projectRepo := repository.NewProjectRepo(pool)
	projectService := usecase.NewProjectService(projectRepo)
	projectHandler := httpHandler.NewProjectHandler(projectService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /task", taskHandler.CreateTask)
	mux.HandleFunc("GET /task/{id}", taskHandler.GetTask)
	mux.HandleFunc("DELETE /task/{id}", taskHandler.DeleteTask)
	mux.HandleFunc("PATCH /task/{id}/comment", taskHandler.AddCommentToTask)
	mux.HandleFunc("PATCH /task/{id}/status", taskHandler.UpdateStatus)

	mux.HandleFunc("POST /register", userHandler.Register)
	mux.HandleFunc("POST /login", userHandler.Login)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUser)

	mux.HandleFunc("POST /project", projectHandler.CreateProject)
	mux.HandleFunc("GET /project/{id}", projectHandler.GetProject)
	mux.HandleFunc("PATCH /project/{id}", projectHandler.AddTask)
	mux.HandleFunc("PATCH /project/{id}/{person}", projectHandler.AddParticipant)
	mux.HandleFunc("DELETE /project/{id}", projectHandler.DeleteProject)

	fmt.Println("Listening port :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Cannot connect to port `:8080` : %s", err)
	}
}
