package api

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Aergiaaa/todolist/handlers"
	"github.com/Aergiaaa/todolist/storage"
)

// Hardcoded templates
var baseTemplate = `<!DOCTYPE html>
<html>
<head>
    <title>Todo List</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 2em; }
        .todo-item { padding: 8px; border-bottom: 1px solid #eee; }
        .completed { text-decoration: line-through; }
    </style>
</head>
<body>
    <h1>Todo List</h1>
    {{template "content" .}}
</body>
</html>`

var listTemplate = `{{define "content"}}
    <form method="post" action="/todos">
        <input type="text" name="title" placeholder="Add new todo">
        <button type="submit">Add</button>
    </form>
    <div class="todo-list">
        {{range .}}
        <div class="todo-item {{if .Completed}}completed{{end}}">
            {{.Title}}
        </div>
        {{else}}
        <p>No todos yet!</p>
        {{end}}
    </div>
{{end}}`

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
	tmpl := template.New("base")
	var err error

	// Parse base template
	tmpl, err = tmpl.Parse(baseTemplate)
	if err != nil {
		log.Println("Failed to parse base template:", err)
		return template.Must(template.New("error").Parse("<html><body>Template error</body></html>"))
	}

	// Parse list template
	_, err = tmpl.Parse(listTemplate)
	if err != nil {
		log.Println("Failed to parse list template:", err)
		return template.Must(template.New("error").Parse("<html><body>Template error</body></html>"))
	}

	return tmpl
}
