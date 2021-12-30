package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/minio/minio-go/v7"
)

type ObjectInfo struct {
	ObjectName string
	BucketSrc  string
	BucketDst  string
	Err        error
}

func moveObjects(minioClient *minio.Client, argsRaw []string) {
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

	chanCopyObject := objectCopier2(minioClient, bucketNameSrc, bucketNameDst, objectCh)
	chanRemoveObject := objectRemover2(minioClient, chanCopyObject)

	counterTotal := 0
	counterSuccess := 0
	for fileResult := range chanRemoveObject {
		if fileResult.Err != nil {
			fmt.Printf("error for %s.\n", fileResult.Err)
		} else {
			counterSuccess++
		}
		counterTotal++
	}

	fmt.Printf("Successfully moving %v/%v objects from %s to %s", counterSuccess, counterTotal, bucketNameSrc, bucketNameDst)

}
func objectCopier2(minioClient *minio.Client, bucketNameSrc string, bucketNameDst string, objectCh <-chan minio.ObjectInfo) <-chan ObjectInfo {
	chanOut := make(chan ObjectInfo)
	go func() {
		for object := range objectCh {
			src := minio.CopySrcOptions{
				Bucket: bucketNameSrc,
				Object: object.Key,
			}
			dst := minio.CopyDestOptions{
				Bucket: bucketNameDst,
				Object: object.Key, //change this to rename copied object
			}
			// Copy object
			fileInfo, err := minioClient.CopyObject(context.Background(), dst, src)
			chanOut <- ObjectInfo{
				ObjectName: fileInfo.Key,
				BucketSrc:  bucketNameSrc,
				BucketDst:  bucketNameDst,
				Err:        err,
			}
		}
		defer close(chanOut)
	}()
	return chanOut
}

func objectRemover2(minioClient *minio.Client, chanIn <-chan ObjectInfo) <-chan ObjectInfo {
	chanOut := make(chan ObjectInfo)

	wg := &sync.WaitGroup{}
	for object := range chanIn {
		wg.Add(1)
		go func(object ObjectInfo) {
			err := objectRemover(minioClient, object.BucketSrc, object.ObjectName)
			chanOut <- ObjectInfo{
				Err: err,
			}
			wg.Done()
		}(object)
	}
	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}
