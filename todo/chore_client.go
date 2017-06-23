package todo

import (
	"fmt"
	"net/http"

	"github.com/benschw/opin-go/rest"
)

// Client Factory
func NewChoreClient(addr string) *ChoreClient {
	return &ChoreClient{
		Addr: addr,
	}
}

// Client
type ChoreClient struct {
	Addr string
}

func (c *ChoreClient) Add(content string, choreType string) (*Chore, error) {
	var created *Chore
	chore := &Chore{Content: content, Type: choreType}

	r, err := rest.MakeRequest("POST", fmt.Sprintf("%s/api/chore", c.Addr), chore)
	if err != nil {
		return created, err
	}
	err = rest.ProcessResponseEntity(r, &created, http.StatusCreated)
	return created, err
}

func (c *ChoreClient) FindAll() ([]*Chore, error) {
	var chores []*Chore

	r, err := rest.MakeRequest("GET", fmt.Sprintf("%s/api/chore", c.Addr), nil)
	if err != nil {
		return chores, err
	}
	err = rest.ProcessResponseEntity(r, &chores, http.StatusOK)
	return chores, err
}

func (c *ChoreClient) Delete(id int) error {
	r, err := rest.MakeRequest("DELETE", fmt.Sprintf("%s/api/chore/%d", c.Addr, id), nil)
	if err != nil {
		return err
	}
	return rest.ProcessResponseEntity(r, nil, http.StatusNoContent)
}
