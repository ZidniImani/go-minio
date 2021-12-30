package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/minio/minio-go/v7"
)

func copyObjects(minioClient *minio.Client, argsRaw []string) {
	ctx, cancel := context.WithCancel(context.Background())
	var bucketNameSrc, bucketNameDst, subFolder string

	// check second arguments
	if len(argsRaw) > 2 {
		bucketNameSrc = argsRaw[1]
		bucketNameDst = argsRaw[2]
	} else {
		panic("Use param: cp <SOURCE_BUCKETNAME> <DESTINATION_BUCKETNAME>")
	}

	defer cancel()

	// list object
	objectCh := minioClient.ListObjects(ctx, bucketNameSrc, minio.ListObjectsOptions{
		Prefix: subFolder + "/",
	})

	wg := &sync.WaitGroup{}
	for object := range objectCh {

		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		// run as goroutine
		go objectCopier(minioClient, bucketNameSrc, bucketNameDst, object.Key, wg)
	}
	wg.Wait()
	fmt.Println(objectCh)
	fmt.Printf("Successfully copy objects from %s to %s", bucketNameSrc, bucketNameDst)
}

func objectCopier(minioClient *minio.Client, bucketNameSrc string, bucketNameDst string, file string, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	src := minio.CopySrcOptions{
		Bucket: bucketNameSrc,
		Object: file,
	}
	dst := minio.CopyDestOptions{
		Bucket: bucketNameDst,
		Object: file, //change this to rename copied object
	}
	// Copy object
	_, err := minioClient.CopyObject(context.Background(), dst, src)
	// fileInfo, err := minioClient.CopyObject(context.Background(), dst, src)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Printf("copy object %s\n", fileInfo.Key)

}
