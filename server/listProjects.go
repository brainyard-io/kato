package server

import (
	pb "github.com/brainyard-io/kato/api"
)

// ListProjects returns a list of project ids
func (s *Server) ListProjects(identifier *pb.Identifier, listProjectsServer pb.Kato_ListProjectsServer) error {
	return nil
}