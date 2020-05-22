package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

// GetUser returns a User from database
func (s *Server) GetUser(ctx context.Context, identifier *pb.Identifier) (*pb.User, error) {
	return &pb.User{}, nil
}