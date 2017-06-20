package todo

import (
	"fmt"
	"net/http"

	"github.com/benschw/opin-go/rest"
)

// Client Factory
func NewTodoClient(addr string) *TodoClient {
	return &TodoClient{
		Addr: addr,
	}
}

// Client
type TodoClient struct {
	Addr string
}

func (c *TodoClient) Add(content string) (*Todo, error) {
	var created *Todo
	todo := &Todo{Content: content, Status: "new"}

	r, err := rest.MakeRequest("POST", fmt.Sprintf("%s/todo", c.Addr), todo)
	if err != nil {
		return created, err
	}
	err = rest.ProcessResponseEntity(r, &created, http.StatusCreated)
	return created, err
}

func (c *TodoClient) Find(id int) (*Todo, error) {
	var found *Todo

	r, err := rest.MakeRequest("GET", fmt.Sprintf("%s/todo/%d", c.Addr, id), nil)
	if err != nil {
		return found, err
	}
	err = rest.ProcessResponseEntity(r, &found, http.StatusOK)
	return found, err
}

func (c *TodoClient) FindAll() ([]*Todo, error) {
	var todos []*Todo

	r, err := rest.MakeRequest("GET", fmt.Sprintf("%s/todo", c.Addr), nil)
	if err != nil {
		return todos, err
	}
	err = rest.ProcessResponseEntity(r, &todos, http.StatusOK)
	return todos, err
}
func (c *TodoClient) Save(todo *Todo) (*Todo, error) {
	var saved *Todo

	r, err := rest.MakeRequest("PUT", fmt.Sprintf("%s/todo/%d", c.Addr, todo.Id), todo)
	if err != nil {
		return saved, err
	}
	err = rest.ProcessResponseEntity(r, &saved, http.StatusOK)
	return saved, err
}
func (c *TodoClient) Delete(id int) error {
	r, err := rest.MakeRequest("DELETE", fmt.Sprintf("%s/todo/%d", c.Addr, id), nil)
	if err != nil {
		return err
	}
	return rest.ProcessResponseEntity(r, nil, http.StatusNoContent)
}
