package server

import (
	"errors"
	"github.com/vds/Restraunt/pkg/database"
)

type Server struct {
	db database.Database
}

func NewServer(database database.Database) (*Server, error) {
	if database == nil {
		return nil, errors.New("server expects a valid database instance")
	}
	return &Server{db:database}, nil
}

func (server *Server) Start(port string) error {
	router, err := NewRouter(server.db)
	if err != nil {
		return nil
	}
	return router.Create(port)
}
