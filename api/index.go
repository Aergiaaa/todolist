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
	todoHandler := handlers.NewTodoHandler(todoStore, templates)

	path := r.URL.Path

	if path == "/" || path == "/todos" || path == "/todos/" {
		todoHandler.ListTodos(w, r)
		return
	}

	http.NotFound(w, r)
}

// loadTemplates loads the hardcoded templates
func loadTemplates() *template.Template {
	tmpl := template.New("")

	// Get template files
	templateFiles, err := filepath.Glob("../templates/*.html")
	if err != nil {
		log.Fatal("Failed to get template files:", err)
	}

	for _, file := range templateFiles {
		log.Println("Loaded template:", file)
	}

	// Parse templates
	tmpl, err = tmpl.ParseFiles(templateFiles...)
	if err != nil {
		log.Fatal("Failed to parse templates:", err)
	}

	return tmpl
}
