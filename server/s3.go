package server

import (
    "github.com/minio/minio-go/v6"
    "log"
)

func (s *Server) initS3Client() *Server {
	// Initialize minio client object.
	var err error
    s.s3Client , err = minio.New(s.args.S3.Endpoint, s.args.S3.AccessKeyID, s.args.S3.SecretAccessKey, s.args.S3.SSL)
    if err != nil {
        log.Fatalln(err)
    }
	log.Printf("%#v\n", s.s3Client) // minioClient is now setup
	return s
}