package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

// PutUser creates/updates a user
func (s *Server) PutUser(ctx context.Context, User *pb.User) (*pb.User, error) {
	return &pb.User{}, nil
}
