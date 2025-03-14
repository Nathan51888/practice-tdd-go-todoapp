package httpserver

import (
	"encoding/json"
	"mytodoapp/domain/todo"
	"net/http"

	"github.com/google/uuid"
)

type Driver struct {
	BaseURL string
	Client  *http.Client
}

func (d Driver) GetTodoById(id uuid.UUID) (todo.Todo, error) {
	res, err := d.Client.Get(d.BaseURL + "/todo?id=" + id.String())
	if err != nil {
		return todo.Todo{}, err
	}
	defer res.Body.Close()

	var result todo.Todo
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return todo.Todo{}, err
	}
	return result, nil
}

func (d Driver) CreateTodo(title string) (todo.Todo, error) {
	res, err := d.Client.Post(d.BaseURL+"/todo?title="+title, "", nil)
	if err != nil {
		return todo.Todo{}, err
	}
	defer res.Body.Close()

	var result todo.Todo
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return todo.Todo{}, err
	}

	return result, nil
}
