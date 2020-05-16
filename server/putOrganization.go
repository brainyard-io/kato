package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

func (s *Server) PutOrganization(ctx context.Context, Organization *pb.Organization) (*pb.Organization, error) {
	return &pb.Organization{}, nil
}