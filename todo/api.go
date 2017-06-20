package todo

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
}

type Task struct {
	Chore      Chore `json:"chore"`
	IsComplete bool  `json:"is-complete"`
	Overdue    int   `json:"overdue"`
}
