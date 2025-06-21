package api

import (
	"embed"
	"html/template"
	"log"
	"net/http"

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

var tmplFS embed.FS

// loadTemplates loads the HTML templates
func loadTemplates() *template.Template {
	tmpl := template.New("")

	// Parse templates from the embedded filesystem
	var err error
	tmpl, err = tmpl.ParseFS(tmplFS, "templates/*.html")
	if err != nil {
		log.Println("Failed to parse templates:", err)
		return template.Must(template.New("error").Parse("<html><body><h1>Template Error</h1></body></html>"))
	}

	log.Println("Templates loaded successfully using embed")
	return tmpl
}
