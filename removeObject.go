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
		panic("Use param: up <BUCKETNAME> <OBJECTNAME>")
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
	} else {
		var input string
		fmt.Printf("Are you sure to remove bucket '%v'? \nWarning: all objects inside bucket will also be deleted (y/n): ", bucketName)
		fmt.Scanln(&input)
		if strings.ToLower(input) == "n" {
			return
		}
	}

	// remove object
	err = minioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Succesfully remove object %s in bucket %s with name %s", objInfo.Key, bucketName, objectName)

}
