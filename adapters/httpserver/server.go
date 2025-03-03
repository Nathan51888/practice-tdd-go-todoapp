package httpserver

import (
	"encoding/json"
	"log"
	"mytodoapp/domain/todo"
	"net/http"
)

type TodoService interface {
	GetTodoByTitle(title string) (todo.Todo, error)
}

type TodoServer struct {
	store todo.TodoStore
	http.Handler
}

func NewTodoServer(store todo.TodoStore) *TodoServer {
	server := new(TodoServer)

	server.store = store

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todo", server.GetTodoByTitle)
	mux.HandleFunc("POST /todo", server.CreateTodo)
	mux.HandleFunc("PUT /todo", server.UpdateTodoTitle)

	server.Handler = mux

	return server
}

func (t *TodoServer) GetTodoByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	result, err := t.store.GetTodoByTitle(title)
	if err != nil {
		log.Printf("Error GetTodoByTitle(): %v", err)
	}

	json.NewEncoder(w).Encode(&result)
}

func (t *TodoServer) CreateTodo(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	result, err := t.store.CreateTodo(title)
	if err != nil {
		log.Printf("Error CreateTodo(): %v", err)
	}

	json.NewEncoder(w).Encode(result)
}

func (t *TodoServer) UpdateTodoTitle(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")
	title := r.URL.Query().Get("title")

	result, err := t.store.UpdateTodoTitle(target, title)
	if err != nil {
		log.Printf("Error UpdateTodoTitle(): %v", err)
	}

	json.NewEncoder(w).Encode(result)
}
