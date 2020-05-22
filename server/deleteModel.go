package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

// DeleteModel deletes a model from the database (not S3)
func (s *Server) DeleteModel(ctx context.Context, model *pb.Model) (*pb.Model, error) {
	return model, nil
}