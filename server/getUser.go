package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

func (s *Server) GetUser(ctx context.Context, identifier *pb.Identifier) (*pb.User, error) {
	return &pb.User{}, nil
}