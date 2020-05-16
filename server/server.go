package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	//"github.com/golang/protobuf/proto"

	pb "github.com/brainyard-io/kato/api"
)

type ServerArgs struct {
	port	string
	address string
	secure	bool
	cert	string
	key		string
	creds	string

	s3AccessKey string
	s3SecretAccessKey string

	databaseHost	string
	databasePort	string
	databaseUser	string
	databaseSecret	string
}

func NewServer(args ServerArgs) *Server {
	s := &Server{
		args: args,
	}
	return s.Init()
}

type Server struct {
	args	ServerArgs
	grpcOpts	[]grpc.ServerOption
	listener	net.Listener
}

func (s *Server) Init() *Server {
	//TODO: Add database
	//TODO: grpc-server
	//TODO: s3 connector
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
	return s
}

func (s *Server) getAddress() string {
	return fmt.Sprintf("%s:%s", s.args.address, s.args.port)
}

func (s *Server) Serve() {
	grpcServer := grpc.NewServer(s.grpcOpts...)
	pb.RegisterKatoServer(grpcServer, s)
	grpcServer.Serve(s.listener)
}