package todo

import (
	"github.com/jinzhu/gorm"
)

type TaskRepo struct {
	ChoreRepo *ChoreRepo
	Db        *gorm.DB
}

func (r *TaskRepo) ToggleStatus(id int) error {
	return nil
}

func (r *TaskRepo) FindAll(choreType string) ([]Task, error) {
	var tasks []Task

	return tasks, nil
}
