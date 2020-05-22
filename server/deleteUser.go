package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

// DeleteUser deletes a user from database
func (s *Server) DeleteUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	return user, nil
}