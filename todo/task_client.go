package todo

import (
	"fmt"
	"net/http"

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

func (c *TaskClient) toggleStatus(id int) (*Task, error) {
	var created *Task

	r, err := rest.MakeRequest("POST", fmt.Sprintf("%s/task/toggle-status/%d", c.Addr, id), nil)
	if err != nil {
		return created, err
	}
	err = rest.ProcessResponseEntity(r, &created, http.StatusCreated)
	return created, err
}

func (c *TaskClient) FindAllDaily() ([]*Task, error) {
	return c.findAllByType(TYPE_DAILY)
}
func (c *TaskClient) FindAllWeekly() ([]*Task, error) {
	return c.findAllByType(TYPE_WEEKLY)
}
func (c *TaskClient) FindAllMonthly() ([]*Task, error) {
	return c.findAllByType(TYPE_MONTHLY)
}
func (c *TaskClient) FindAllYearly() ([]*Task, error) {
	return c.findAllByType(TYPE_YEARLY)
}

func (c *TaskClient) findAllByType(choreType string) ([]*Task, error) {
	var tasks []*Task

	r, err := rest.MakeRequest("GET", fmt.Sprintf("%s/task/%s", c.Addr, choreType), nil)
	if err != nil {
		return tasks, err
	}
	err = rest.ProcessResponseEntity(r, &tasks, http.StatusOK)
	return tasks, err
}
