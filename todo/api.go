package todo

import (
	"time"
)

const (
	TYPE_DAILY   = "daily"
	TYPE_WEEKLY  = "weekly"
	TYPE_MONTHLY = "monthly"
	TYPE_YEARLY  = "yearly"
)

type Chore struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Tasks   []Task `json:"tasks";"gorm:"ForeignKey:ChoreId"`
}

type Task struct {
	Id      int       `json:"id"`
	ChoreId int       `json:"chore-id"`
	Time    time.Time `json:"time"`
}
