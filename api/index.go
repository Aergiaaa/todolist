package api

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Aergiaaa/todolist/handlers"
	"github.com/Aergiaaa/todolist/storage"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Initialize the storage
	todoStore := storage.NewMemoryStorage()

	// Load templates
	templates := loadTemplates()

	// Initialize the handlers
	todoHandler := handlers.NewTodoHandler(todoStore, templates) // Set up routes

	path := r.URL.Path

	if path == "/" {
		// Directly serve the todos list on the home page
		todoHandler.ListTodos(w, r)
		return
	}

	if path == "/todos" || path == "/todos/" {
		todoHandler.ListTodos(w, r)
		return
	}

	if len(path) > 8 && path[:8] == "/static/" {
		// For static files, we'll need to handle differently in serverless
		// This is simplified and might need adjustment
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
		return
	}

	http.NotFound(w, r)
}

// loadTemplates loads the HTML templates
func loadTemplates() *template.Template {
	// Create a new template with empty name
	tmpl := template.New("")

	// Get template files
	templateFiles, err := filepath.Glob("templates/*.html")
	if err != nil {
		log.Println("Failed to get template files:", err)
	}

	// Parse templates
	tmpl, err = tmpl.ParseFiles(templateFiles...)
	if err != nil {
		log.Println("Failed to parse templates:", err)
	}

	return tmpl
}
