package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/minio/minio-go/v7"
)

func removeObject(minioClient *minio.Client, argsRaw []string) {
	var bucketName, objectName string

	// check arguments
	if len(argsRaw) > 2 {
		bucketName = argsRaw[1]
		objectName = argsRaw[2]
	} else {
		panic("Use param: rm <BUCKETNAME> <OBJECTNAME>")
	}

	// check bucket
	_, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// check object
	objInfo, err := minioClient.StatObject(context.Background(), bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		fmt.Println("Requested object is not exists")
		return
	}

	// confirm
	var input string
	fmt.Printf("Are you sure to remove '%v'? \n(y/n): ", bucketName)
	fmt.Scanln(&input)
	if strings.ToLower(input) == "n" {
		return
	}

	// remove object
	err = objectRemover(minioClient, bucketName, objectName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// display result
	fmt.Printf("Succesfully remove object %s from bucket '%s'\n", objInfo.Key, bucketName)
}

// for reusable function
func objectRemover(minioClient *minio.Client, bucketName string, objectName string) error {
	return minioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
}
