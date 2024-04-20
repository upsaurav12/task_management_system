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

func main() {

	r := mux.NewRouter()

	tasks = append(tasks, Task{ID: "1", Title: "First Title", Description: "First Description"})
	tasks = append(tasks, Task{ID: "2", Title: "Second Title", Description: "Second Description"})
	tasks = append(tasks, Task{ID: "3", Title: "Third Title", Description: "Third Description"})
	tasks = append(tasks, Task{ID: "4", Title: "Fourth Title", Description: "Fourth Description"})
	tasks = append(tasks, Task{ID: "5", Title: "Fifth Title", Description: "Fifth Description"})
	tasks = append(tasks, Task{ID: "6", Title: "Sixth Title", Description: "Sixth Description"})

	r.HandleFunc("/api/tasks", getTasks).Methods("GET")
	r.HandleFunc("/api/tasks/{id}", getTask).Methods("GET")
	r.HandleFunc("/api/tasks", createTask).Methods("POST")
	r.HandleFunc("/api/tasks/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/api/tasks/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
