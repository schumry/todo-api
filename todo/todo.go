package todo

import (
	"math/rand"
	"time"
)

type Todo struct {
	Value     string
	Completed bool
	Id        int64
	Time      string
}

func NewTodos() Todos {
	return Todos{
		Todos: make([]Todo, 0),
	}
}

type Todos struct {
	Todos []Todo
}

func (t *Todos) Get() []Todo {
	return t.Todos
}

func (t *Todos) Add(todo Todo) Todo {
	todo.Id = int64(rand.Int31())
	todo.Time = time.Now().Format(time.RFC3339)
	t.Todos = append(t.Todos, todo)
	return todo
}

func (t *Todos) Delete(id int64) {
	var index int = -1

	for i, todo := range t.Todos {
		if todo.Id == id {
			index = i
		}
	}

	if index != -1 {
		t.Todos[index] = t.Todos[len(t.Todos)-1]
		t.Todos[len(t.Todos)-1] = Todo{}
		t.Todos = t.Todos[:len(t.Todos)-1]
	} else {
		return
	}
}

func (t *Todos) Complete(id int64) {
	var index int = -1

	for i, todo := range t.Todos {
		if todo.Id == id {
			index = i
		}
	}

	if index != -1 {
		t.Todos[index].Completed = true
	} else {
		return
	}
}
