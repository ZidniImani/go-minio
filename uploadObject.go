package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
)

func uploadObject(minioClient *minio.Client, argsRaw []string) {
	var bucketName, storagePath, objectName string

	// check arguments
	if len(argsRaw) > 2 {
		bucketName = argsRaw[1]
		storagePath = argsRaw[2]
	} else {
		panic("Use param: up <BUCKETNAME> <FILEPATH>")
	}

	// check bucket
	_, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Upload
	objectName = storagePath[strings.LastIndex(storagePath, `\`)+1:]
	newFileName := "upload-" + objectName

	// open file from this computer
	file, err := os.Open(storagePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	// upload/put object
	_, err = minioClient.PutObject(context.Background(),
		bucketName,
		newFileName,
		file,
		fileStat.Size(),
		minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Successfully uploaded %s of size %d\n", objectName, fileStat.Size())

}
