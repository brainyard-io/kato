package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

func (s *Server) GetOrganization(ctx context.Context, identifier *pb.Identifier) (*pb.Organization, error) {
	return &pb.Organization{}, nil
}