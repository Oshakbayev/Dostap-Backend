package bucket

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type EventBucketServiceInterface interface {
	UploadFile(string, string, multipart.File) error
}

func (b *Bucket) UploadFile(bucketName string, objectKey string, file multipart.File) error {
	_, err := b.bucketClient.PutObject(context.TODO(), &s3.PutObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(objectKey),
        Body:   file,
    })
    return err
}