package main

import (
	"fmt"
	"os"
)

func argumentReader(argsRaw []string) {
	// start connection
	minioClient := GetMinIOConnection()

	// bucket operation
	switch {
	case argsRaw[0] == "ls":
		getList(minioClient, argsRaw)
	case argsRaw[0] == "mb":
		createBucket(minioClient, argsRaw) //default policy built-in
	case argsRaw[0] == "rb":
		removeBucket(minioClient, argsRaw)
	// object operation
	case argsRaw[0] == "up":
		uploadObject(minioClient, argsRaw)
	case argsRaw[0] == "rm":
		removeObject(minioClient, argsRaw)
	case argsRaw[0] == "cp": //copy all object
		copyObjects(minioClient, argsRaw)
	case argsRaw[0] == "mv":
		moveObjects(minioClient, argsRaw)
	// set/get policy
	case argsRaw[0] == "policy":
		bucketPolicy(minioClient, argsRaw)
	default:
		panic("Unknown argument")
	}
}

// unknown arguments handler
func unknownArgs() {
	if err := recover(); err != nil {
		fmt.Printf("Error: %v", err)
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
