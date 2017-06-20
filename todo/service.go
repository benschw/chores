package todo

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/benschw/opin-go/ophttp"
	"github.com/gorilla/mux"
)

func NewService(bind string, conStr string) (*TodoService, error) {
	server := ophttp.NewServer(bind)

	db, err := gorm.Open("mysql", conStr)
	if err != nil {
		return nil, err
	}
	return &TodoService{
		Server: server,
		Db:     db,
	}, nil
}

type TodoService struct {
	Server *ophttp.Server
	Db     *gorm.DB
}

// Migrate
func (s *TodoService) Migrate() {
	s.Db.AutoMigrate(&Chore{})
	s.Db.AutoMigrate(&Task{})
}

// Configure and start http server
func (s *TodoService) Run() error {
	defer s.Db.Close()

	// Build Resource
	choreRepo := &ChoreRepo{Db: s.Db}
	chores := &ChoreResource{Repo: choreRepo}

	taskRepo := &TaskRepo{ChoreRepo: choreRepo, Db: s.Db}
	tasks := &TaskResource{Repo: taskRepo}

	// Wire Routes
	r := mux.NewRouter()

	r.HandleFunc("/chore", chores.Add).Methods("POST")
	r.HandleFunc("/chore", chores.GetAll).Methods("GET")
	r.HandleFunc("/chore/{id}", chores.Delete).Methods("DELETE")

	r.HandleFunc("/task/toggle-status/{id}", tasks.ToggleStatus).Methods("POST")
	r.HandleFunc("/task/daily", tasks.GetAllDaily).Methods("GET")
	r.HandleFunc("/task/weekly", tasks.GetAllWeekly).Methods("GET")
	r.HandleFunc("/task/monthly", tasks.GetAllMonthly).Methods("GET")
	r.HandleFunc("/task/yearly", tasks.GetAllYearly).Methods("GET")

	mux := http.NewServeMux()
	mux.Handle("/", r)

	// Start Server
	err := s.Server.Start(mux)

	log.Println("Server Stopped")
	return err
}

func (s *TodoService) Stop() {
	log.Println("Stopping Server...")
	s.Server.Stop()
}
