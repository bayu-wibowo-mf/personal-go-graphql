package main

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	ctx := context.Background()
	endpoint := "localhost:9000"
	accessKeyID := "acFBcDcF3isclvRIG04S"
	secretAccessKey := "MnOHyuO6QoDVUYNjA5ORNu4WQSp9stC26ezOgQMe"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})

	if err != nil {
		log.Fatal(err)
	}

	objectName := "test-object"
	filePath := "./test-file.txt"

	err = minioClient.FGetObject(ctx, "graphql-prototype", objectName, filePath, minio.GetObjectOptions{})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File downloaded successfully to %s\n", filePath)
}
