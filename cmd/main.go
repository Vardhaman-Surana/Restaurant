package main

import (
	"github.com/vds/Restraunt/pkg/database/mysql"
	"github.com/vds/Restraunt/pkg/server"
)

func main() {

	// create database instance
	db, err := mysql.NewMySqlDB()
	if err != nil {
		panic(err)
	}

	// create server
	s, err := server.NewServer(db)
	if err != nil {
		panic(err)
	}
	if err := s.Start(":5000"); err != nil {
		panic(err)
	}
}
