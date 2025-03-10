package inmemory

import (
	"mytodoapp/domain/todo"
)

type InMemoryTodoStore struct {
	Todos []todo.Todo
}

func NewInMemoryTodoStore() (*InMemoryTodoStore, error) {
	return &InMemoryTodoStore{}, nil
}

func (i *InMemoryTodoStore) GetTodoAll() ([]todo.Todo, error) {
	return i.Todos, nil
}

func (i *InMemoryTodoStore) GetTodoByTitle(title string) (todo.Todo, error) {
	var result todo.Todo
	for _, todo := range i.Todos {
		if todo.Title == title {
			result = todo
		}
	}
	return result, nil
}

func (i *InMemoryTodoStore) GetTodoById(todoId int) (todo.Todo, error) {
	var result todo.Todo
	for _, item := range i.Todos {
		if item.Id == todoId {
			result = item
		}
	}
	return result, nil
}

func (i *InMemoryTodoStore) CreateTodo(title string) (todo.Todo, error) {
	createdTodo := todo.Todo{Title: title, Completed: false}
	createdTodo.Id = len(i.Todos) + 1
	i.Todos = append(i.Todos, createdTodo)
	return createdTodo, nil
}

func (i *InMemoryTodoStore) UpdateTodoTitle(todoId int, title string) (todo.Todo, error) {
	var result todo.Todo
	for index, todo := range i.Todos {
		if todo.Id == todoId {
			i.Todos[index].Title = title
			result = i.Todos[index]
		}
	}
	return result, nil
}

func (i *InMemoryTodoStore) UpdateTodoStatus(todoId int, completed bool) (todo.Todo, error) {
	var result todo.Todo
	for index, todo := range i.Todos {
		if todo.Id == todoId {
			i.Todos[index].Completed = completed
			result = i.Todos[index]
		}
	}
	return result, nil
}

func (i *InMemoryTodoStore) UpdateTodoById(todoId int, changedTodo todo.Todo) (todo.Todo, error) {
	var result todo.Todo
	for index, item := range i.Todos {
		if item.Id == todoId {
			i.Todos[index] = changedTodo
			result = i.Todos[index]
		}
	}
	return result, nil
}
