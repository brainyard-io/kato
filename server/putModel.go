package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

// PutModel creates/updates a model
func (s *Server) PutModel(ctx context.Context, model *pb.Model) (*pb.Model, error) {
	return &pb.Model{}, nil
}