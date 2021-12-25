package main

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func getBucketList(minioClient *minio.Client, ctx context.Context) {
	buckets, err := minioClient.ListBuckets(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bucket := range buckets {
		fmt.Println(bucket.Name)
	}
}
