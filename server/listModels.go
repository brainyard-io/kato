package server

import (
	pb "github.com/brainyard-io/kato/api"
)

// ListModels returns a list of model ids
func (s *Server) ListModels(identifier *pb.Identifier, listModelsServer pb.Kato_ListModelsServer) error {
	return nil
}
