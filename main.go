package main

import (
	"context"
	"fmt"
	"os"
)

func argumentReader(argsRaw []string) {
	// start connection
	minioClient := GetMinIOConnection()
	ctx := context.Background()

	// policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::hellobucket/*"],"Sid": ""}]}`

	// err := minioClient.SetBucketPolicy(context.Background(), "hellobucket", policy)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// policy, err := minioClient.GetBucketPolicy(context.Background(), "hellobucket2")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("policy %s\n", policy)

	// read arguments
	switch {
	case argsRaw[0] == "ls":
		getList(minioClient, ctx, argsRaw)
	case argsRaw[0] == "mb":
		createBucket(minioClient, ctx, argsRaw)
	case argsRaw[0] == "rb":
		removeBucket(minioClient, ctx, argsRaw)
	// object operation
	case argsRaw[0] == "up":
		uploadObject(minioClient, ctx, argsRaw)
	case argsRaw[0] == "rm":
		removeObject(minioClient, ctx, argsRaw)
	// policy
	case argsRaw[0] == "policy":
		bucketPolicy(minioClient, ctx, argsRaw)
	default:
		panic("Unknown argument")
	}
}

func unknownArgs() {
	if err := recover(); err != nil {
		fmt.Println("Unknown argument parameter")
		fmt.Printf("Error! %v", err)
	}
}

func main() {
	// read argument
	var argsRaw = os.Args[1:]

	// catch unknown arg
	defer unknownArgs()

	// call arg reader
	argumentReader(argsRaw)
}
