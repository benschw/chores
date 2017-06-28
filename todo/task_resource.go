package todo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/benschw/opin-go/rest"
)

type TaskResource struct {
	Repo *TaskRepo
}

func (r *TaskResource) LogWork(res http.ResponseWriter, req *http.Request) {
	var task Task

	if err := rest.Bind(req, &task); err != nil {
		log.Print(err)
		rest.SetBadRequestResponse(res)
		return
	}

	task, err := r.Repo.LogWork(task)
	if err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}

	if err := rest.SetCreatedResponse(res, task, fmt.Sprintf("task/%d", task.Id)); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
}
func (r *TaskResource) DeleteWork(res http.ResponseWriter, req *http.Request) {
	id, err := rest.PathInt(req, "id")
	if err != nil {
		rest.SetBadRequestResponse(res)
		return
	}

	if err := r.Repo.DeleteWork(id); err != nil {
		rest.SetBadRequestResponse(res)
		return
	}

	if err := rest.SetNoContentResponse(res); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
}
func (r *TaskResource) GetAll(res http.ResponseWriter, req *http.Request) {
	tasks, err := r.Repo.FindAll()
	if err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}

	if err := rest.SetOKResponse(res, tasks); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
}
