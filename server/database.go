package server

import (
    "github.com/go-pg/pg"
    //"github.com/go-pg/pg/orm"
)

// Initiate the postgres Client
func (s *Server) initDatabaseClient() *Server {
	s.db = pg.Connect(&pg.Options{
		User: s.args.Database.User,
		Addr: s.args.Database.Host,
		Password: s.args.Database.Secret,
		Database: s.args.Database.Database,
	})
	return s
}