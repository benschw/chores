package main

import (
	"flag"
	"log"
	"os"

	"github.com/benschw/chores/todo"
)

func main() {

	bind := flag.String("bind", "0.0.0.0:80", "address to bind http server to")
	conStr := flag.String("mysql", "admin:changeme@tcp(localhost:3306)/Chores?charset=utf8&parseTime=True", "db connection string")
	flag.Parse()

	db, found := os.LookupEnv("MYSQL_CONNECTION_STRING")
	if found {
		conStr = &db
	}

	log.Print("constructing service")
	svc, err := todo.NewService(*bind, *conStr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Print("migrating")
	svc.Migrate()

	log.Print("running service")
	if err := svc.Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
