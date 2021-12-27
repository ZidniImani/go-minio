package main

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func getList(minioClient *minio.Client, ctx context.Context, argsRaw []string) {

	// check arguments
	switch {
	case len(argsRaw) >= 2:
		listBucketObject(minioClient, argsRaw)
	case len(argsRaw) == 1:
		listBucket(minioClient, ctx)
	default:
		panic("Argument not valid")
	}
}

func listBucket(minioClient *minio.Client, ctx context.Context) {
	buckets, err := minioClient.ListBuckets(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bucket := range buckets {
		fmt.Println(bucket.Name)
	}
}
