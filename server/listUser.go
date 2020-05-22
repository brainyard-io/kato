package server

import (
	pb "github.com/brainyard-io/kato/api"
)

// ListUser returns a list of user ids
func (s *Server) ListUser(identifier *pb.Identifier, listUserServer pb.Kato_ListUserServer) error {
	return nil
}