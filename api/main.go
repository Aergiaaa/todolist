package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Aergiaaa/todolist/handlers"
	"github.com/Aergiaaa/todolist/storage"
)

func main() {
	// Initialize the storage
	todoStore := storage.NewMemoryStorage()

	// Load templates
	templates := loadTemplates()

	// Initialize the handlers
	todoHandler := handlers.NewTodoHandler(todoStore, templates) // Set up routes
	http.Handle("/todos/", todoHandler)
	http.Handle("/todos", todoHandler) // Add a route without the trailing slash
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		// Directly serve the todos list on the home page
		todoHandler.ListTodos(w, r)
	})

	// Start the server
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// loadTemplates loads the HTML templates
func loadTemplates() *template.Template {
	// Create a new template with empty name
	tmpl := template.New("")

	// Get template files
	templateFiles, err := filepath.Glob("templates/*.html")
	if err != nil {
		log.Fatal("Failed to get template files:", err)
	}

	// Parse templates
	tmpl, err = tmpl.ParseFiles(templateFiles...)
	if err != nil {
		log.Fatal("Failed to parse templates:", err)
	}

	return tmpl
}
