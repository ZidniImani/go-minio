# MinIO Client - Golang
A simple MinIO client app build with Go.

S3 object storage: [MinIO](https://min.io/).

# Features
1. Bucket operation: create, remove, list
2. Object operation: upload, remove, copy, move
3. Set/get bucket policy

# How to Run
Make sure you have [MinIO server](https://docs.min.io/docs/minio-quickstart-guide.html) that can be connected and have [Go](https://go.dev/doc/install) in your computer. To run, you can run it directly with Go or build it first.
## Set connection
This project use [MinIO server](https://docs.min.io/docs/minio-quickstart-guide.html) on localhost, you can change it on connection.go if needed.
## Run directly with Go
```go
go run . <your_command>
```

# Command list
## Bucket
1. Create
    ```go
    go run . mb <bucketname>
    ```
    
2. Remove
    ```go
    go run . rb <bucketname>
    ```
    
3. List all bucket
    ```go
    go run . ls
    ```
    
4. Set policy
    ```go
    go run . policy <bucketname> <your_policy.yaml>
    ```
    See more about policy on [official site](https://docs.min.io/minio/baremetal/security/minio-identity-management/policy-based-access-control.html) and [set policy on Go](https://docs.min.io/docs/golang-client-api-reference.html#SetBucketPolicy) or you can use example default.yaml
    
5. Get policy
    ```go
    go run . policy <bucketname>
    ```
## Object
1. Upload an object
    ```go
    go run . up <bucketname> <yourfilepath>
    ```
    
2. List object from bucket
    ```go
    go run . ls <bucketname>
    ```
    Use / for see object in subfolder. Example: go run . ls mybucket/mysubfolder.
    
3. Remove an object
    ```go
    go run . rm <bucketname> <yourobject>
    ```
    
3. Copy all object in bucket to another bucket
    ```go
    go run . cp <bucket_source> <bucket_destination>
    ```
    
4. Move all object in bucket to another bucket
    ```go
    go run . mv <bucket_source> <bucket_destination>
    ```
    
## How to build and run as executable
```go
go build .
```
then run as executable in terminal/shell
```sh
./go-minio.exe <your_command> 
```