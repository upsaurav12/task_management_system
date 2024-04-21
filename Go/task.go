package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var tasks []Task

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {

}

func createTask(w http.ResponseWriter, r *http.Request) {

}

func updateTask(w http.ResponseWriter, r *http.Request) {

}
func deleteTask(w http.ResponseWriter, r *http.Request) {

}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			// Handle preflight request
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r) // Continue with request handling
	})
}

func main() {
	r := mux.NewRouter()
	r.Use(corsMiddleware) // Apply the CORS middleware globally

	// Define your routes...

	// Pre-fill some tasks for testing
	tasks = append(tasks, Task{ID: "1", Title: "First Task", Description: "First Description"})
	tasks = append(tasks, Task{ID: "2", Title: "Second Task", Description: "Second Description"})
	tasks = append(tasks, Task{ID: "3", Title: "Third Task", Description: "Third Description"})
	tasks = append(tasks, Task{ID: "4", Title: "Fourth Task", Description: "Fifth Description"})

	// Handle CORS preflight requests
	// Define routes for the RESTful API
	r.HandleFunc("/api/tasks", getTasks).Methods("GET")
	r.HandleFunc("/api/tasks/{id}", getTask).Methods("GET")
	r.HandleFunc("/api/tasks", createTask).Methods("POST")
	r.HandleFunc("/api/tasks/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/api/tasks/{id}", deleteTask).Methods("DELETE")

	// Start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", r))

}
