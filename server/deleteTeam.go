package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

func (s *Server) DeleteTeam(ctx context.Context, team *pb.Team) (*pb.Team, error) {
	return team, nil
}