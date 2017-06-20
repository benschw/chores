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

func GetClientAndService() (*TodoClient, *TodoService) {

	dbStr := "admin:changeme@tcp(172.20.20.1:3306)/Todo?charset=utf8&parseTime=True"
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

	client := &TodoClient{
		Addr: fmt.Sprintf("http://localhost:%d", port),
	}
	return client, svc
}
