package server

import (
	pb "github.com/brainyard-io/kato/api"
)

// ListOrganizations returns a list of organization ids
func (s *Server) ListOrganizations(identifier *pb.Identifier, listOrganizationsServer pb.Kato_ListOrganizationsServer) error {
	return nil
}
