package todo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TodoStoreContract struct {
	NewTodoStore func() (TodoStore, error)
}

func (c TodoStoreContract) Test(t *testing.T) {
	t.Run("can get all todos as slice from database", func(t *testing.T) {
		sut, err := c.NewTodoStore()
		if err != nil {
			t.Fatalf("Error creating todo store: %v\n", err)
		}

		want := []Todo{
			{Title: "Todo1", Completed: false},
			{Title: "Todo2", Completed: false},
			{Title: "Todo3", Completed: false},
		}
		sut.CreateTodo("Todo1")
		sut.CreateTodo("Todo2")
		sut.CreateTodo("Todo3")

		got, err := sut.GetTodoAll()
		assert.NoError(t, err)
		for index, item := range got {
			if item.Title != want[index].Title {
				t.Fatal("title not equal")
			}
			if item.Completed != want[index].Completed {
				t.Fatal("completed not equal")
			}
		}
	})
	t.Run("can create, get, update todo's title and status by title from database", func(t *testing.T) {
		sut, err := c.NewTodoStore()
		if err != nil {
			t.Fatalf("Error creating todo store: %v\n", err)
		}

		want := Todo{Title: "Todo_new", Completed: false}
		newTodo, err := sut.CreateTodo("Todo_new")
		assert.NoError(t, err)
		assert.Equal(t, want.Title, newTodo.Title)
		assert.Equal(t, want.Completed, newTodo.Completed)
		got, err := sut.GetTodoByTitle("Todo_new")
		assert.NoError(t, err)
		assert.Equal(t, newTodo, got)

		want = Todo{Id: got.Id, Title: "Todo_updated", Completed: got.Completed}
		updatedTodo, err := sut.UpdateTodoTitle(want.Id, "Todo_updated")
		assert.NoError(t, err)
		assert.Equal(t, want, updatedTodo)
		got, err = sut.GetTodoByTitle("Todo_updated")
		assert.NoError(t, err)
		assert.Equal(t, want, got)

		want = Todo{Id: got.Id, Title: "Todo_updated", Completed: true}
		updatedTodo, err = sut.UpdateTodoStatus(want.Id, true)
		assert.NoError(t, err)
		assert.Equal(t, want, updatedTodo)
		got, err = sut.GetTodoByTitle("Todo_updated")
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
	t.Run("can update todo by id", func(t *testing.T) {
		sut, err := c.NewTodoStore()
		require.NoError(t, err)

		want := Todo{Title: "Todo_new", Completed: false}
		newTodo, err := sut.CreateTodo("Todo_new")
		assert.NoError(t, err)
		assert.Equal(t, want.Title, newTodo.Title)
		assert.Equal(t, want.Completed, newTodo.Completed)
		got, err := sut.GetTodoById(newTodo.Id)
		assert.NoError(t, err)
		assert.Equal(t, newTodo, got, "GetTodoById()")

		want = Todo{Id: got.Id, Title: "Todo_updated", Completed: true}
		updatedTodo, err := sut.UpdateTodoById(want.Id, want)
		assert.NoError(t, err)
		assert.Equal(t, want, updatedTodo, "UpdateTodoById()")
		got, err = sut.GetTodoById(updatedTodo.Id)
		assert.NoError(t, err)
		assert.Equal(t, want, got, "GetTodoById()")
	})
	t.Run("can delete todo by id", func(t *testing.T) {
		sut, err := c.NewTodoStore()
		require.NoError(t, err)

		// TODO: dry it with function
		want := Todo{Title: "Delete_this", Completed: false}
		newTodo, err := sut.CreateTodo("Delete_this")
		assert.NoError(t, err)
		assert.Equal(t, want.Title, newTodo.Title)
		assert.Equal(t, want.Completed, newTodo.Completed)
		got, err := sut.GetTodoById(newTodo.Id)
		assert.NoError(t, err)
		assert.Equal(t, newTodo, got, "GetTodoById()")

		want = Todo{Id: got.Id, Title: "Delete_this", Completed: false}
		deletedTodo, err := sut.DeleteTodoById(want.Id)
		assert.NoError(t, err, "DeleteTodoById()")
		assert.Equal(t, want, deletedTodo, "DeleteTodoById()")
		got, err = sut.GetTodoById(want.Id)
		assert.Error(t, err, "GetTodoById()")
		assert.NotEqual(t, want, got, "GetTodoById()")
	})
}
