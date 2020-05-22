package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

// GetModel returns a model with s3 download link
func (s *Server) GetModel(ctx context.Context, identifier *pb.Identifier) (*pb.Model, error) {
	return &pb.Model{}, nil
}