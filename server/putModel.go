package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

func (s *Server) PutModel(ctx context.Context, model *pb.Model) (*pb.Model, error) {
	return &pb.Model{}, nil
}