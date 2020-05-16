package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

func (s *Server) GetModel(ctx context.Context, identifier *pb.Identifier) (*pb.Model, error) {
	return &pb.Model{}, nil
}