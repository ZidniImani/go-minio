package main

import (
    "log"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
    endpoint := "127.0.0.1:57040/"
    accessKeyID := "secretminioadmin"
    secretAccessKey := "secretminioadmin"
    useSSL := true

    // Initialize minio client object.
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
        Secure: useSSL,
    })
    if err != nil {
        log.Fatalln("error!")
        log.Fatalln(err)
    }

    log.Printf("Success") // minioClient is now set up
    log.Printf("%#v\n", minioClient)
}
