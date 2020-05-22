package server

import (
	"context"

	pb "github.com/brainyard-io/kato/api"
)

// Deletes an organization from database
func (s *Server) DeleteOrganization(ctx context.Context, organization *pb.Organization) (*pb.Organization, error) {
	return organization, nil
}