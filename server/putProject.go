package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

// PutProject creates/updates a project
func (s *Server) PutProject(ctx context.Context, Project *pb.Project) (*pb.Project, error) {
	return &pb.Project{}, nil
}
