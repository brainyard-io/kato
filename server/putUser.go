package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

func (s *Server) PutUser(ctx context.Context, User *pb.User) (*pb.User, error) {
	return &pb.User{}, nil
}