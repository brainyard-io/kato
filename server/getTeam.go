package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

// Returns a team from database
func (s *Server) GetTeam(ctx context.Context, identifier *pb.Identifier) (*pb.Team, error) {
	return &pb.Team{}, nil
}