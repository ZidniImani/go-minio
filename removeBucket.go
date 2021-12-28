package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/minio/minio-go/v7"
)

func removeBucket(minioClient *minio.Client, argsRaw []string) {
	var bucketName string

	// check second arguments
	if len(argsRaw) > 1 {
		bucketName = argsRaw[1]
	} else {
		panic("Bucket name not found in params")
	}

	// check bucket
	found, err := minioClient.BucketExists(context.Background(), bucketName)
	// bucket not found
	if err != nil {
		fmt.Println(err)
		return
	}
	// warn user
	if found {
		var input string
		fmt.Printf("Are you sure to remove bucket '%v'? \nWarning: all objects inside bucket will also be deleted (y/n): ", bucketName)
		fmt.Scanln(&input)
		if strings.ToLower(input) == "n" {
			return
		}
	}

	// remove bucket
	err = minioClient.RemoveBucket(context.Background(), bucketName)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Successfully remove bucket %s\n", bucketName)
	}

}
