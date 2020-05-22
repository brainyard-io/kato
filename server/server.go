package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	//"github.com/golang/protobuf/proto"

	pb "github.com/brainyard-io/kato/api"
	pg "github.com/go-pg/pg"
	minio "github.com/minio/minio-go/v6"
)

// ServerArgs represents the arguments kato server starts with
type ServerArgs struct {
	port    string
	address string
	secure  bool
	cert    string
	key     string
	creds   string
	S3      struct {
		AccessKeyID     string
		SecretAccessKey string
		Endpoint        string
		SSL             bool
	}
	Database struct {
		Host     string
		User     string
		Secret   string
		Database string
	}
}

// NewServer returns a new server instance
func NewServer(args ServerArgs) *Server {
	s := &Server{
		args: args,
	}
	return s.Init()
}

// Server structure
type Server struct {
	args     ServerArgs
	grpcOpts []grpc.ServerOption
	listener net.Listener
	s3       *minio.Client
	db       *pg.DB
}

// Init initiates the server
func (s *Server) Init() *Server {
	//TODO: Add database
	//TODO: grpc-server
	//TODO: nats-eventing
	var err error
	s.listener, err = net.Listen("tcp", s.getAddress())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if s.args.secure {
		if s.args.cert == "" {
			log.Fatalf("No Certificate specified")
		}
		if s.args.key == "" {
			log.Fatalf("No Key specified")
		}
		creds, err := credentials.NewServerTLSFromFile(s.args.cert, s.args.key)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		s.grpcOpts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	s.initS3Client()
	s.initDatabaseClient()
	s.initPrometheusEndpoint()
	return s
}

// getAddress returns an address from host and port
func (s *Server) getAddress() string {
	return fmt.Sprintf("%s:%s", s.args.address, s.args.port)
}

// Serve starts to serve grpc
func (s *Server) Serve() {
	grpcServer := grpc.NewServer(s.grpcOpts...)
	pb.RegisterKatoServer(grpcServer, s)
	grpcServer.Serve(s.listener)
}
