package awsS3

import (
	"context"
	"fmt"

	"lambda-sample/pkg/model"
	"lambda-sample/pkg/wUtil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	awsRegion   = "aws_region"
	awsS3Bucket = "aws_s3_bucket"
)

func GetServerLog(ctx context.Context, file string) ([]byte, error) {
	bucket := os.Getenv("aws_s3_server_log")
	if len(bucket) == 0 {
		return nil, fmt.Errorf("bucket empty")
	}
	awsSess, err := getAWSSession()
	if err != nil {
		return []byte{}, wUtil.NewError(err)
	}

	buf := aws.NewWriteAtBuffer([]byte{})
	downloader := s3manager.NewDownloader(awsSess)
	_, err = downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(file),
	})
	if err != nil {
		return []byte{}, wUtil.NewError(err)
	}
	return buf.Bytes(), nil
}

func LoadKSFile(ctx aws.Context, fileKey string) ([]byte, error) {
	awsSess, err := getAWSSession()
	if err != nil {
		return []byte{}, wUtil.NewError(err)
	}

	bucket := os.Getenv("aws_s3_ks")
	if len(bucket) == 0 {
		return nil, wUtil.NewError(model.ErrS3BucketEmpty)
	}
	buf := aws.NewWriteAtBuffer([]byte{})
	downloader := s3manager.NewDownloader(awsSess)
	_, err = downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileKey),
	})
	if err != nil {
		err = fmt.Errorf("%v : %v", err, bucket)
		return []byte{}, wUtil.NewError(err)
	}

	return buf.Bytes(), nil
}
