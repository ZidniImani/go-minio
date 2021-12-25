package main

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func listBucketObject(minioClient *minio.Client, argsRaw []string) {
	ctx, cancel := context.WithCancel(context.Background())
	var bucketName string

	// check second arguments
	if len(argsRaw) > 1 {
		bucketName = argsRaw[1]
	} else {
		panic("Bucket name not found in params")
	}

	defer cancel()

	objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		// Prefix:    "myprefix",
		// Recursive: true,
	})

	fmt.Printf("Object from bucket %v:\n", bucketName)

	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println("-", object.Key)
	}
}
