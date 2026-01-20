package main

import (
	"context"
	"fmt"
	"log"
	"mime"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func main() {
	localFilePath := os.Args[1]
	rustfsEndpoint := "http://115.190.54.31:8082"
	accessKey := "w9OaC2EYPgIFu7QqpKHy"
	secretKey := "4cik8GmyTJn79vHNtKw0YquLZ2ehrOpWXBD5zgMl"
	bucketName := "rui"
	objectKey := filepath.Base(localFilePath)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(_, _ string, _ ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: rustfsEndpoint}, nil
			})),
	)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}
	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	file, err := os.Open(localFilePath)
	if err != nil {
		log.Fatalf("打开文件失败: %v", err)
	}
	defer file.Close()
	contentType := mime.TypeByExtension(filepath.Ext(localFilePath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectKey),
		Body:        file,
		ContentType: aws.String(contentType),
		ACL:         types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		log.Fatalf("上传失败: %v", err)
	}
	fmt.Printf("文件上传成功\n")
	fmt.Printf("访问地址：%s/%s/%s\n", rustfsEndpoint, bucketName, objectKey)
}
