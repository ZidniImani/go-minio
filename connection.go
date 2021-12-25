package main

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioClient *minio.Client
)

func initiateMinioConnection() *minio.Client {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "secretminioadmin"
	secretAccessKey := "secretminioadmin"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}

func GetMinIOConnection() *minio.Client {
	if minioClient == nil {
		minioClient = initiateMinioConnection()
	}
	return minioClient
}
