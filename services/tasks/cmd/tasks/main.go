package main

import (
	"log"
	"net/http"
	"os"

	tasksauth "example.com/tech-ip-sem2/services/tasks/internal/client/authclient"
	taskshttp "example.com/tech-ip-sem2/services/tasks/internal/http"
	"example.com/tech-ip-sem2/services/tasks/internal/service"
	"example.com/tech-ip-sem2/shared/middleware"
)

func main() {
	port := os.Getenv("TASKS_PORT")
	if port == "" {
		port = "8082"
	}

	authBaseURL := os.Getenv("AUTH_BASE_URL")
	if authBaseURL == "" {
		authBaseURL = "http://localhost:8081"
	}

	mux := http.NewServeMux()
	handler := taskshttp.NewHandler(service.NewTaskService(), tasksauth.New(authBaseURL))
	handler.Register(mux)

	wrapped := middleware.RequestID(middleware.Logging("tasks")(mux))

	addr := ":" + port
	log.Printf("tasks service started on %s, auth=%s", addr, authBaseURL)
	if err := http.ListenAndServe(addr, wrapped); err != nil {
		log.Fatal(err)
	}
}
