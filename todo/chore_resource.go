package todo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/benschw/opin-go/rest"
)

type ChoreResource struct {
	Repo *ChoreRepo
}

func (r *ChoreResource) Add(res http.ResponseWriter, req *http.Request) {
	var chore Chore

	if err := rest.Bind(req, &chore); err != nil {
		log.Print(err)
		rest.SetBadRequestResponse(res)
		return
	}

	chore, err := r.Repo.Insert(chore)
	if err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}

	if err := rest.SetCreatedResponse(res, chore, fmt.Sprintf("chore/%d", chore.Id)); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
}

func (r *ChoreResource) GetAll(res http.ResponseWriter, req *http.Request) {
	chores, err := r.Repo.FindAll()
	if err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}

	if err := rest.SetOKResponse(res, chores); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
}

func (r *ChoreResource) Delete(res http.ResponseWriter, req *http.Request) {
	id, err := rest.PathInt(req, "id")
	if err != nil {
		rest.SetBadRequestResponse(res)
		return
	}

	if err := r.Repo.Delete(id); err != nil {
		rest.SetBadRequestResponse(res)
		return
	}

	if err := rest.SetNoContentResponse(res); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
}
