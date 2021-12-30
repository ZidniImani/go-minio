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
		panic("Use param: mv <SOURCE_BUCKETNAME> <DESTINATION_BUCKETNAME>")
	}

	defer cancel()

	// list object
	objectCh := minioClient.ListObjects(ctx, bucketNameSrc, minio.ListObjectsOptions{
		Prefix: subFolder + "/",
	})

	// pipeline
	chanCopyObject := objectCopier2(minioClient, bucketNameSrc, bucketNameDst, objectCh)
	chanRemoveObject := objectRemover2(minioClient, chanCopyObject)

	//counting result
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

	// display result
	fmt.Printf("Successfully moving %v/%v objects from %s to %s\n", counterSuccess, counterTotal, bucketNameSrc, bucketNameDst)

}
func objectCopier2(minioClient *minio.Client, bucketNameSrc string, bucketNameDst string, objectCh <-chan minio.ObjectInfo) <-chan ObjectInfo {
	// create channel
	chanOut := make(chan ObjectInfo)
	// run as goroutine
	go func() {
		for object := range objectCh {
			// set source
			src := minio.CopySrcOptions{
				Bucket: bucketNameSrc,
				Object: object.Key,
			}
			// set destination
			dst := minio.CopyDestOptions{
				Bucket: bucketNameDst,
				Object: object.Key, //change this to rename copied object
			}
			// Copy object
			fileInfo, err := minioClient.CopyObject(context.Background(), dst, src)
			// input to channel
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
	// for every item in channel
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
	// dispatch goroutine to wait all
	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}
