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

	r, err := rest.MakeRequest("POST", fmt.Sprintf("%s/work/", c.Addr), task)
	if err != nil {
		return created, err
	}
	err = rest.ProcessResponseEntity(r, &created, http.StatusCreated)
	return created, err
}
func (c *TaskClient) DeleteWork(id int) error {
	r, err := rest.MakeRequest("DELETE", fmt.Sprintf("%s/work/%d", c.Addr, id), nil)
	if err != nil {
		return err
	}
	return rest.ProcessResponseEntity(r, nil, http.StatusNoContent)
}

func (c *TaskClient) FindAllDaily() ([]*Chore, error) {
	return c.findAllByType(TYPE_DAILY)
}
func (c *TaskClient) FindAllWeekly() ([]*Chore, error) {
	return c.findAllByType(TYPE_WEEKLY)
}
func (c *TaskClient) FindAllMonthly() ([]*Chore, error) {
	return c.findAllByType(TYPE_MONTHLY)
}
func (c *TaskClient) FindAllYearly() ([]*Chore, error) {
	return c.findAllByType(TYPE_YEARLY)
}

func (c *TaskClient) findAllByType(choreType string) ([]*Chore, error) {
	var chores []*Chore

	r, err := rest.MakeRequest("GET", fmt.Sprintf("%s/task/%s", c.Addr, choreType), nil)
	if err != nil {
		return chores, err
	}
	err = rest.ProcessResponseEntity(r, &chores, http.StatusOK)
	return chores, err
}