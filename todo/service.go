package todo

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/benschw/opin-go/ophttp"
	"github.com/gorilla/handlers"
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
func (r *TodoService) Health(res http.ResponseWriter, req *http.Request) {
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
	r.HandleFunc("/", s.Health).Methods("POST")

	r.HandleFunc("/api/chore", chores.Add).Methods("POST")
	r.HandleFunc("/api/chore/{id}", chores.Save).Methods("PUT")
	r.HandleFunc("/api/chore", chores.GetAll).Methods("GET")
	r.HandleFunc("/api/chore/{id}", chores.Delete).Methods("DELETE")

	r.HandleFunc("/api/work", tasks.LogWork).Methods("POST")
	r.HandleFunc("/api/work/{id}", tasks.DeleteWork).Methods("DELETE")

	r.HandleFunc("/api/task", tasks.GetAll).Methods("GET")

	mux := http.NewServeMux()
	mux.Handle("/", handlers.LoggingHandler(os.Stdout, handlers.CORS()(r)))

	// Start Server
	err := s.Server.Start(mux)

	log.Println("Server Stopped")
	return err
}

func (s *TodoService) Stop() {
	log.Println("Stopping Server...")
	s.Server.Stop()
}
