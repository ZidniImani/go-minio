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

func createBucket(minioClient *minio.Client, ctx context.Context, argsRaw []string) {
	var bucketName, policy string

	// check second arguments
	if len(argsRaw) > 1 {
		bucketName = argsRaw[1]
	} else {
		panic("Bucket name not found in params")
	}

	// create bucket
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check bucket
		bucketExists, err := minioClient.BucketExists(ctx, bucketName)
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
	err = minioClient.SetBucketPolicy(context.Background(), bucketName, policy)
	if err != nil {
		fmt.Printf("Successfully created bucket %s but fail to set policy\n", bucketName)
	} else {
		fmt.Printf("Successfully create bucket %s and set policy", bucketName)
	}

}

func readPolicy(yamlPath string) string {
	var policy string
	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		fmt.Printf("Failed to read policy file.\n%+v", err)
		return ""
	}
	var config *Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("Failed to unmarshal: %+v", err)
		return ""
	}
	policy = strings.Fields(config.Policy)[0]
	return policy
}
