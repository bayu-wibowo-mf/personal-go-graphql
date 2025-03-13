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
	contentType := "text/plain"

	info, err := minioClient.FPutObject(ctx, "graphql-prototype", objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v\n", info)
}
