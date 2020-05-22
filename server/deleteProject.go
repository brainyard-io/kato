package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

// Deletes a project from database
func (s *Server) DeleteProject(ctx context.Context, project *pb.Project) (*pb.Project, error) {
	return project, nil
}