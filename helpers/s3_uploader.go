package helpers

import (
	"bytes"
	"context"
	"mime/multipart"

	"github.com/akhil-is-watching/techletics_alumni_reg/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Uploader struct {
	S3Client *s3.Client
}

var uploader *Uploader

func InitS3Uploader() {

	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	creds := credentials.NewStaticCredentialsProvider(config.AwsAccessKeyId, config.AwsSecretAccessKey, "")

	client := s3.NewFromConfig(aws.Config{
		BaseEndpoint: aws.String("https://s3.ap-south-1.amazonaws.com"),
		Region:       *aws.String(config.AwsRegion),
		Credentials:  creds,
	})

	uploader = &Uploader{S3Client: client}
}

func GetS3Uploader() *Uploader {
	return uploader
}

func (uploader *Uploader) Upload(objectKey string, file multipart.File) error {
	_, err := uploader.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("techleticsassetbucket"),
		Key:    aws.String(objectKey),
		Body:   file,
	})

	if err != nil {
		return err
	}

	return nil
}

func (uploader *Uploader) UploadBytes(objectKey string, file []byte) error {
	_, err := uploader.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("techleticsassetbucket"),
		Key:    aws.String(objectKey),
		Body:   bytes.NewReader(file),
	})

	if err != nil {
		return err
	}

	return nil
}
