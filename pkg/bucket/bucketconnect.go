package bucket

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

func ConnectToBucket() *s3.Client {
	//cfg, err := config.LoadDefaultConfig(context.TODO())
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-north-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("AKIA6ODUZQEEMNBOLB52", "/+Z7U8VRCGLko6Xb0RLnoneB9IHhNIpnzpAzmNUo", "")),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := s3.NewFromConfig(cfg)

	return client
}
