package todo

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type ChoreRepo struct {
	Db *gorm.DB
}

func (r *ChoreRepo) Insert(chore Chore) (Chore, error) {
	chore.Created = time.Now()
	r.Db.Create(&chore)

	return chore, nil
}

func (r *ChoreRepo) Update(chore Chore) (Chore, error) {
	chore.Created = time.Now()
	r.Db.Save(&chore)

	return chore, nil
}
func (r *ChoreRepo) Find(id int) (Chore, error) {
	var chore Chore

	if r.Db.First(&chore, id).RecordNotFound() {
		return chore, fmt.Errorf("Not Found")
	}
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
