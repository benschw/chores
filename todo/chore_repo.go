package todo

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type ChoreRepo struct {
	Db *gorm.DB
}

func (r *ChoreRepo) Insert(chore Chore) (Chore, error) {

	r.Db.Create(&chore)

	return chore, nil
}

func (r *ChoreRepo) FindAll() ([]Chore, error) {
	var chores []Chore

	r.Db.Find(&chores)

	return chores, nil
}

func (r *ChoreRepo) Delete(id int) error {
	var chore Chore

	if r.Db.First(&chore, id).RecordNotFound() {
		return fmt.Errorf("chore not found")
	}
	r.Db.Delete(&chore)
	return nil
}
