package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

func (s *Server) DeleteModel(ctx context.Context, model *pb.Model) (*pb.Model, error) {
	return model, nil
}