package todo

import (
	"fmt"
	"net/http"
	"time"

	"github.com/benschw/opin-go/rest"
)

// Client Factory
func NewTaskClient(addr string) *TaskClient {
	return &TaskClient{
		Addr: addr,
	}
}

// Client
type TaskClient struct {
	Addr string
}

func (c *TaskClient) LogWork(choreId int, t time.Time) (*Task, error) {
	task := &Task{ChoreId: choreId, Time: t}
	var created *Task

	r, err := rest.MakeRequest("POST", fmt.Sprintf("%s/api/work/", c.Addr), task)
	if err != nil {
		return created, err
	}
	err = rest.ProcessResponseEntity(r, &created, http.StatusCreated)
	return created, err
}
func (c *TaskClient) DeleteWork(id int) error {
	r, err := rest.MakeRequest("DELETE", fmt.Sprintf("%s/api/work/%d", c.Addr, id), nil)
	if err != nil {
		return err
	}
	return rest.ProcessResponseEntity(r, nil, http.StatusNoContent)
}

func (c *TaskClient) FindAll() ([]*Chore, error) {
	var chores []*Chore

	r, err := rest.MakeRequest("GET", fmt.Sprintf("%s/api/task", c.Addr), nil)
	if err != nil {
		return chores, err
	}
	err = rest.ProcessResponseEntity(r, &chores, http.StatusOK)
	return chores, err
}
