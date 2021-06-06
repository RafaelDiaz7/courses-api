package main

import (
	"courses-api-mysql-and-cb/internal/server"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	host = "localhost"
	port = 9500
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
