package main

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func createBucket(minioClient *minio.Client, ctx context.Context, argsRaw []string) {
	var bucketName string

	// check second arguments
	if len(argsRaw) > 1 {
		bucketName = argsRaw[1]
	} else {
		panic("Bucket name not found in params")
	}

	// create bucket
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			fmt.Printf("Bucket %s is exist\n", bucketName)
		} else {
			fmt.Print(err)
		}
	} else {
		fmt.Printf("Successfully created bucket %s\n", bucketName)
	}

}
