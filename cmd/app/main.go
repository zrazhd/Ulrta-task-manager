package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	httpHandler "github.com/zrazhd/Ulrta-task-manager/internal/delivery/http"
	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
	repository "github.com/zrazhd/Ulrta-task-manager/internal/repository/postgres"
	cache "github.com/zrazhd/Ulrta-task-manager/internal/repository/redis"
	"github.com/zrazhd/Ulrta-task-manager/internal/usecase"
)

func main() {
	ctx := context.Background()
	connStr := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Cannot connect to Database: %s", err)
	}

	taskCache := cache.NewCache[domain.Task](client, "task", time.Duration(time.Second*60))
	userCache := cache.NewCache[domain.User](client, "user", time.Duration(time.Second*60))
	projectCache := cache.NewCache[domain.Project](client, "project", time.Duration(time.Second*60))

	taskRepo := repository.NewTaskRepo(pool)
	taskService := usecase.NewTaskService(taskRepo, taskCache)
	taskHandler := httpHandler.NewTaskHandler(taskService)

	userRepo := repository.NewUserRepo(pool)
	userService := usecase.NewUserService(userRepo, userCache)
	userHandler := httpHandler.NewUserHandler(userService)

	projectRepo := repository.NewProjectRepo(pool)
	projectService := usecase.NewProjectService(projectRepo, projectCache)
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
