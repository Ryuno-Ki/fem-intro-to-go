package main

import (
  "fmt"
  "log"
  "net/http"
  "html/template"
)

// Properties must match HTML template!
type Todo struct {
  Title string
  Content string
}

// Represents context for template
type PageVariables struct {
  PageTitle string
  PageTodos []Todo
}

// Go built-in types as required by function handlers
func home (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Home!")
}

func getTodo (w http.ResponseWriter, r *http.Request) {
  // Respond with plain text
  fmt.Fprint(w, "Todos!")
}

// Fake a database
var todos []Todo

func getTodos (w http.ResponseWriter, r *http.Request) {
  // Properties must match template!
  pageVariables := PageVariables{
    PageTitle: "Get Todos",
    PageTodos: todos,
  }

  // Parsed template
  t, err := template.ParseFiles("todos.html")

  if err != nil {
    // Return a 400 Bad Request with the error message
    http.Error(w, err.Error(), http.StatusBadRequest)
    log.Print("Template parsing error: ", err)
  }

  // Second argument is context, pass nil if omit
  err = t.Execute(w, pageVariables)
}

func addTodo (w http.ResponseWriter, r *http.Request) {
  // Unpack HTTP Form that got sent over the wire
  err := r.ParseForm()

  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    log.Print("Request parsing error: ", err)
  }

  todo := Todo{
    // Input name as argument
    Title: r.FormValue("title"),
    Content: r.FormValue("content"),
  }

  todos = append(todos, todo)
  log.Print(todos)

  // This is not best practice
  http.Redirect(w, r, "/todos/", http.StatusSeeOther)
}

func main () {
	// First argument is path, second argument is callback
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	http.HandleFunc("/add-todo/", addTodo)
	http.HandleFunc("/todo/", getTodo)

	fmt.Println("Server is running on port :8080")
	// Second argument is kind of optional
	// The ListenAndServe is always executed.
	// If something goes wrong, it will return an error which then gets
	// handled by log.Fatal()
	// Important! Run Go from root of project!
	log.Fatal(http.ListenAndServe(":8080", nil))
}
