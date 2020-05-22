package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

// Creates/Updates an organization
func (s *Server) PutOrganization(ctx context.Context, Organization *pb.Organization) (*pb.Organization, error) {
	return &pb.Organization{}, nil
}