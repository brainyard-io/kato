package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

// PutOrganization creates/updates an organization
func (s *Server) PutOrganization(ctx context.Context, Organization *pb.Organization) (*pb.Organization, error) {
	return &pb.Organization{}, nil
}
