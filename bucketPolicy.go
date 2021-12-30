package main

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func bucketPolicy(minioClient *minio.Client, argsRaw []string) {
	var bucketName, policy string
	// check second arguments
	if len(argsRaw) > 1 {
		bucketName = argsRaw[1]
	} else {
		panic("Bucket name not found in params")
	}
	if len(argsRaw) > 2 {
		// set policy
		policy = fmt.Sprintf(readPolicy(argsRaw[2]), bucketName)
		err := minioClient.SetBucketPolicy(context.Background(), bucketName, policy)
		if err != nil {
			fmt.Println("Fail to set policy")
			return
		}
		fmt.Printf("Successfully set bucket %s's policy\n", bucketName)
	} else {
		// get policy
		policy, err := minioClient.GetBucketPolicy(context.Background(), bucketName)
		if err != nil {
			fmt.Printf("Failed to get policy bucket '%s': %s\n", bucketName, err)
			return
		}
		fmt.Printf("Bucket policy: %s\n", policy)
	}
}
