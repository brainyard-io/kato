package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

// PutTeam creates/updates a team
func (s *Server) PutTeam(ctx context.Context, Team *pb.Team) (*pb.Team, error) {
	return &pb.Team{}, nil
}
