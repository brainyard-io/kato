package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

// Returns a project
func (s *Server) GetProject(ctx context.Context, identifier *pb.Identifier) (*pb.Project, error) {
	return &pb.Project{}, nil
}