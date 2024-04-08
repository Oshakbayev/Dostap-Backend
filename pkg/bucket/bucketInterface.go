package bucket

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type BucketInterface interface {
	EventBucketServiceInterface
}

type Bucket struct {
	log          *log.Logger
	bucketClient *s3.Client
}

func CreateBucket(log *log.Logger, bucketClient *s3.Client) BucketInterface {
	return &Bucket{log: log, bucketClient: bucketClient}
}
