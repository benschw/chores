package todo

import (
	"fmt"
	"testing"

	"github.com/benschw/opin-go/ophttp"
	"github.com/benschw/opin-go/rando"
	"github.com/jinzhu/gorm"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&TestSuite{})

type TestSuite struct {
	svc    *TodoService
	tasks  *TaskClient
	chores *ChoreClient
}

func (s *TestSuite) SetUpSuite(c *C) {
	s.chores, s.tasks, s.svc = GetClientAndService()

	go s.svc.Run()
}
func (s *TestSuite) TearDownSuite(c *C) {
	s.svc.Stop()
}
func (s *TestSuite) SetUpTest(c *C) {
	s.svc.Migrate()
}
func (s *TestSuite) TearDownTest(c *C) {
	s.svc.Db.DropTable(Chore{})
	s.svc.Db.DropTable(Task{})
}

func GetClientAndService() (*ChoreClient, *TaskClient, *TodoService) {
	//travis
	//	dbStr := "root:@tcp(127.0.0.1:3306)/Chores?charset=utf8&parseTime=True"

	//linux
	//dbStr := "admin:changeme@tcp(172.17.0.1:3306)/Chores?charset=utf8&parseTime=True"

	//mac
	dbStr := "admin:changeme@tcp(172.20.20.1:3306)/Chores?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", dbStr)
	if err != nil {
		panic(err)
	}

	port := rando.Port()
	bind := fmt.Sprintf("0.0.0.0:%d", port)
	server := ophttp.NewServer(bind)

	svc := &TodoService{
		Server: server,
		Db:     db,
	}

	chores := &ChoreClient{
		Addr: fmt.Sprintf("http://localhost:%d", port),
	}
	tasks := &TaskClient{
		Addr: fmt.Sprintf("http://localhost:%d", port),
	}
	return chores, tasks, svc
}
