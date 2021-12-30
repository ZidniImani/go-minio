package main

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func getList(minioClient *minio.Client, argsRaw []string) {

	// check arguments
	switch {
	case len(argsRaw) >= 2: // go to object list
		listBucketObject(minioClient, argsRaw)
	case len(argsRaw) == 1:
		listBucket(minioClient)
	default:
		panic("Argument not valid")
	}
}

func listBucket(minioClient *minio.Client) {
	// get bucket list
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	// display bucket
	for _, bucket := range buckets {
		fmt.Println("-", bucket.Name)
	}
}
