# Todo List Application

## Description

A modern, responsive Todo List application built with Go and HTMX. This application allows users to manage their tasks with an intuitive interface and real-time updates without requiring a full page reload.

## Technology Used

- **Backend**: Go (Golang) - Pure Go implementation with no external dependencies
- **Frontend**: HTML, CSS, HTMX
- **UI Framework**: HTMX for dynamic content updates without JavaScript
- **Storage**: In-memory (with architecture that allows easy extension to databases)
- **Architecture**: Clean separation of concerns (MVC-like pattern)

## Features

- **CRUD Operations**:
  - Create new todo items with title and description
  - Read/view the list of todos
  - Update existing todo items
  - Delete todo items
- **Dynamic UI Updates**:
  - Real-time content updates using HTMX
  - No page reloads required for actions
- **Responsive Design**:
  - Works on mobile, tablet and desktop devices
- **Multiple Views**:
  - List view - displays all todos
  - Form view - for creating and editing todos
- **State Management**:
  - Toggle completion status of todos
  - Visual indication of completed items

## Setup Instructions

1. **Prerequisites**:

   - Go 1.18 or higher
   - Web browser

2. **Installation**:

   ```bash
   # Clone the repository
   git clone https://github.com/Aergiaaa/todolist.git

   # Navigate to the project directory
   cd todolist
   ```

3. **Running the Application**:

   ```bash
   # Start the server
   go run main.go
   ```

4. **Access the Application**:
   - Open your web browser
   - Visit http://localhost:8080
   - Start adding and managing your todos!

## Project Structure

```
todolist/
├── handlers/        # HTTP request handlers
├── models/          # Data models
├── storage/         # Storage implementations
├── static/          # Static assets (CSS, JS)
│   ├── css/
│   └── js/
├── templates/       # HTML templates
└── main.go          # Application entry point
```

## AI Support

This project was developed with the assistance of IBM Granite, an AI programming assistant. IBM Granite helped with:

- Initial project structure design
- Generating boilerplate Go code for handlers and storage
- Creating HTML templates with HTMX integration
- Designing responsive CSS
- Implementing CRUD operations
- Documentation generation

The AI helped accelerate development while maintaining clean code practices and modern application architecture. Human supervision and direction were provided throughout the process to ensure the application met specific requirements and maintained high quality standards.
