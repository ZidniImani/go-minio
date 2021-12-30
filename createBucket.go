package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/minio/minio-go/v7"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Policy string `yaml:"policy"`
}

func createBucket(minioClient *minio.Client, argsRaw []string) {
	var bucketName, policy string

	// check second arguments
	if len(argsRaw) > 1 {
		bucketName = argsRaw[1]
	} else {
		panic("Bucket name not found in params")
	}

	// create bucket
	err := minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check bucket
		bucketExists, err := minioClient.BucketExists(context.Background(), bucketName)
		if err == nil && bucketExists {
			fmt.Printf("Bucket %s is exist\n", bucketName)
			return
		} else {
			fmt.Print(err)
			return
		}
	}
	// create policy
	if len(argsRaw) > 2 {
		policy = fmt.Sprintf(readPolicy(argsRaw[2]), bucketName)
	} else {
		policy = fmt.Sprintf(readPolicy("default.yaml"), bucketName)
	}
	// set policy
	err = minioClient.SetBucketPolicy(context.Background(), bucketName, policy)
	if err != nil {
		fmt.Printf("Successfully created bucket %s but fail to set policy\n", bucketName)
		return
	}

	// display result
	fmt.Printf("Successfully create bucket %s and set policy\n", bucketName)

}

func readPolicy(yamlPath string) string {
	var policy string
	// read yaml file
	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		fmt.Printf("Failed to read policy file.\n%+v", err)
		return ""
	}
	var config *Config
	// unmarshal yaml
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("Failed to unmarshal: %+v\n", err)
		return ""
	}
	// convert to string
	policy = strings.Fields(config.Policy)[0]
	return policy
}
