package todo

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type TaskRepo struct {
	ChoreRepo *ChoreRepo
	Db        *gorm.DB
}

func (r *TaskRepo) LogWork(task Task) (Task, error) {
	r.Db.Create(&task)
	return task, nil
}
func (r *TaskRepo) DeleteWork(id int) error {
	var task Task

	if r.Db.First(&task, id).RecordNotFound() {
		return fmt.Errorf("task not found")
	}
	r.Db.Delete(&task)
	return nil
}

func (r *TaskRepo) FindAll(choreType string) ([]Chore, error) {
	var chores []Chore

	r.Db.Preload("Tasks").Where("type = ?", choreType).Find(&chores)

	return chores, nil
}
