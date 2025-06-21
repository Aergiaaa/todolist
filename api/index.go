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
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovered from panic:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

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
		log.Println("Serving static files from:", path)

		http.Error(w, "Static files are not implemented yet", http.StatusNotImplemented)
		return
	}

	http.NotFound(w, r)
}

// loadTemplates loads the HTML templates
func loadTemplates() *template.Template {
	// Create a new template with empty name
	tmpl := template.New("")

	// Get template files
	templateFiles, err := filepath.Glob("./templates/*.html")
	if err != nil {
		log.Println("Failed to get template files:", err)
		return template.Must(template.New("error").Parse("Template loading error"))
	}

	// Parse templates
	tmpl, err = tmpl.ParseFiles(templateFiles...)
	if err != nil {
		log.Println("Failed to parse templates:", err)
		return template.Must(template.New("error").Parse("Template parsing error"))
	}

	return tmpl
}
