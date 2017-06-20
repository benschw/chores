package todo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/benschw/opin-go/rest"
	"github.com/jinzhu/gorm"
)

type TodoResource struct {
	Db *gorm.DB
}

func (r *TodoResource) Health(res http.ResponseWriter, req *http.Request) {
	//2xx => pass, 429 => warn, anything else => critical

	var todos []Todo
	assoc := r.Db.Find(&todos)
	if assoc.Error != nil {
		rest.SetInternalServerErrorResponse(res, nil)
		return
	}
	// set health to OK
	if err := rest.SetOKResponse(res, nil); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
}

func (r *TodoResource) Add(res http.ResponseWriter, req *http.Request) {
	var todo Todo

	if err := rest.Bind(req, &todo); err != nil {
		log.Print(err)
		rest.SetBadRequestResponse(res)
		return
	}

	r.Db.Create(&todo)

	if err := rest.SetCreatedResponse(res, todo, fmt.Sprintf("todo/%d", todo.Id)); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
}

func (r *TodoResource) Get(res http.ResponseWriter, req *http.Request) {
	id, err := rest.PathInt(req, "id")
	if err != nil {
		rest.SetBadRequestResponse(res)
		return
	}
	var todo Todo

	if r.Db.First(&todo, id).RecordNotFound() {
		rest.SetNotFoundResponse(res)
		return
	}

	if err := rest.SetOKResponse(res, todo); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
}

func (r *TodoResource) GetAll(res http.ResponseWriter, req *http.Request) {
	var todos []Todo

	r.Db.Find(&todos)

	if err := rest.SetOKResponse(res, todos); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
}
func (r *TodoResource) Update(res http.ResponseWriter, req *http.Request) {
	var todo Todo

	id, err := rest.PathInt(req, "id")
	if err != nil {
		rest.SetBadRequestResponse(res)
		return
	}
	if err := rest.Bind(req, &todo); err != nil {
		rest.SetBadRequestResponse(res)
		return
	}
	todo.Id = id

	var found Todo
	if r.Db.First(&found, id).RecordNotFound() {
		rest.SetNotFoundResponse(res)
		return
	}

	r.Db.Save(&todo)

	if err := rest.SetOKResponse(res, todo); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}

}
func (r *TodoResource) Delete(res http.ResponseWriter, req *http.Request) {
	id, err := rest.PathInt(req, "id")
	if err != nil {
		rest.SetBadRequestResponse(res)
		return
	}
	var todo Todo

	if r.Db.First(&todo, id).RecordNotFound() {
		rest.SetNotFoundResponse(res)
		return
	}

	r.Db.Delete(&todo)

	if err := rest.SetNoContentResponse(res); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
}
