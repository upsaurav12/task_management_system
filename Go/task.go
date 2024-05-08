package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

var tasks []Task

func connect() (*sql.DB, error) {
	dsn := "saurav:SA@2003up_@tcp(localhost:3306)/tasks"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}

func getTasks(w http.ResponseWriter, r *http.Request) {

	db, err := connect()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := "SELECT id, title, description, priority FROM tasks"

	rows, err := db.Query(query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	/*vars := mux.Vars(r)  // Get URL variables
	taskID := vars["id"] // Get the task ID from the URL

	db, err := connect()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Query to get the task by ID
	query := "SELECT id, title, description FROM tasks WHERE id = ?"
	row := db.QueryRow(query, taskID) // Query for a single result

	var task Task
	if err := row.Scan(&task.ID, &task.Title, &task.Description); err != nil {
		if err == sql.ErrNoRows {
			// If no rows found, return a 404 status
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			// Otherwise, return a server error
			http.Error(w, "Database query error", http.StatusInternalServerError)
		}
		return
	}

	// Set the response content type and return the task as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task) // Return the task data*/
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db, err := connect()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	query := "INSERT INTO tasks (title , description , priority) VALUES(?, ?, ?)"

	result, err := db.Exec(query, task.Title, task.Description, task.Priority)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	taskID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "ERROR", http.StatusInternalServerError)
		return
	}
	task.ID = int(taskID)

	w.Header().Set("Content-Type", "appliation/json")
	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	var task Task
	if err := json.NewDecoder(r.Body).Decode((&task)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := connect()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	query := "UPDATE tasks SET title = ?, description = ? WHERE id = ?"

	_, err = db.Exec(query, task.Title, task.Description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fmt.Sprintf("Task %s updated", taskID))

}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	db, err := connect()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	query := "DELETE FROM tasks WHERE id = ?"

	_, err = db.Exec(query, taskID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fmt.Sprintf("Task %s is deleted ", taskID))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			// Handle preflight request
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r) // Continue with request handling
	})
}

func main() {

	/*db, err := connect()

	if err != nil {
		log.Fatal("Failed to connect mysql:", err)
	} else {
		fmt.Println("Connected")
	}
	defer db.Close()*/

	r := mux.NewRouter()
	r.Use(corsMiddleware) // Apply the CORS middleware globally

	// Define your routes...

	// Pre-fill some tasks for testing
	/*tasks = append(tasks, Task{ID: "1", Title: "First Task", Description: "First Description"})
	tasks = append(tasks, Task{ID: "2", Title: "Second Task", Description: "Second Description"})
	tasks = append(tasks, Task{ID: "3", Title: "Third Task", Description: "Third Description"})
	tasks = append(tasks, Task{ID: "4", Title: "Fourth Task", Description: "Fifth Description"})*/

	// Handle CORS preflight requests
	// Define routes for the RESTful API
	r.HandleFunc("/api/tasks", getTasks).Methods("GET")
	r.HandleFunc("/api/tasks/{id}", getTask).Methods("GET")
	r.HandleFunc("/api/tasks", createTask).Methods("POST")
	r.HandleFunc("/api/tasks/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/api/tasks/{id}", deleteTask).Methods("DELETE")
	r.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		// Handle preflight `OPTIONS` requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusNoContent) // 204 No Content
	}).Methods("OPTIONS")

	r.HandleFunc("/api/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Handle preflight `OPTIONS` requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusNoContent) // 204 No Content
	}).Methods("OPTIONS")

	// Start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", r))

}
