package server

import (
	"github.com/minio/minio-go/v6"
	"log"
)

// initS3Client initiates a S3 client based on minio
func (s *Server) initS3Client() *Server {
	// Initialize minio client object.
	var err error
	s.s3, err = minio.New(s.args.S3.Endpoint, s.args.S3.AccessKeyID, s.args.S3.SecretAccessKey, s.args.S3.SSL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%#v\n", s.s3) // minioClient is now setup
	return s
}
