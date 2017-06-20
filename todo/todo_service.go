package todo

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/benschw/opin-go/ophttp"
	"github.com/gorilla/mux"
)

func NewTodoService(bind string, conStr string) (*TodoService, error) {
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
	s.Db.AutoMigrate(&Todo{})
}

// Configure and start http server
func (s *TodoService) Run() error {
	defer s.Db.Close()

	// Build Resource
	resource := &TodoResource{Db: s.Db}

	// Wire Routes
	r := mux.NewRouter()
	r.HandleFunc("/health", resource.Health).Methods("GET")
	r.HandleFunc("/todo", resource.Add).Methods("POST")
	r.HandleFunc("/todo", resource.GetAll).Methods("GET")
	r.HandleFunc("/todo/{id}", resource.Get).Methods("GET")
	r.HandleFunc("/todo/{id}", resource.Update).Methods("PUT")
	r.HandleFunc("/todo/{id}", resource.Delete).Methods("DELETE")

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
