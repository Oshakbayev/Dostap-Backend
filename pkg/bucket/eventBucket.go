package bucket

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type EventBucketServiceInterface interface {
	UploadFile(string, string, multipart.File) error
	ListObjects(string, string) ([]string, error)
}

func (b *Bucket) ListObjects(bucketName string, prefix string) ([]string, error) {
	result, err := b.bucketClient.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		Prefix: &prefix,
	})
	var contents []string
	if err != nil {
		b.log.Printf("Couldn't list objects in bucket %v. Here's why: %v\n", bucketName, err)
	} else {
		for _, v := range result.Contents {
			contents = append(contents, *v.Key)
		}
	}
	return contents, err
}

func (b *Bucket) UploadFile(bucketName string, objectKey string, file multipart.File) error {
	_, err := b.bucketClient.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		b.log.Printf("error during insert into bucket in UploadFile(repo):%s", err.Error())
	}
	return err
}
