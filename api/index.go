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
	log.Println("Storage initialized")

	// Load templates
	log.Println("Loading templates...")
	templates := loadTemplates()
	log.Println("Templates loaded")

	// Initialize the handlers
	log.Println("Setting up handlers...")
	todoHandler := handlers.NewTodoHandler(todoStore, templates) // Set up routes
	log.Println("Handlers set up")

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

	// Try different paths that might work in Vercel
	templatePaths := []string{
		"./templates/*.html",
		"/templates/*.html",
		"../templates/*.html",
	}

	var templateFiles []string
	var err error

	// Try each path until we find templates
	for _, path := range templatePaths {
		templateFiles, err = filepath.Glob(path)
		if err == nil && len(templateFiles) > 0 {
			log.Println("Found templates at:", path)
			break
		}
	}

	if len(templateFiles) == 0 {
		log.Println("No template files found")
		return template.Must(template.New("error").Parse("<html><body><h1>Todo List</h1><p>Template loading error</p></body></html>"))
	}

	// Parse templates
	tmpl, err = tmpl.ParseFiles(templateFiles...)
	if err != nil {
		log.Println("Failed to parse templates:", err)
		return template.Must(template.New("error").Parse("<html><body><h1>Todo List</h1><p>Template parsing error</p></body></html>"))
	}

	return tmpl
}
