package awsS3

import (
	"bytes"
	"context"
	"fmt"
	"lambda-sample/pkg/model"
	"lambda-sample/pkg/wUtil"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	awsRegion   = "aws_region"
	awsS3Bucket = "aws_s3_bucket"
)

func GenPreSignAvatarURL(ctx context.Context, customerID int) (string, error) {
	avatarBucket := os.Getenv("aws_s3_avatar")
	if len(avatarBucket) == 0 {
		return "", wUtil.ErrorWithStr("aws_s3_avatar empty")
	}
	sess, err := getAWSSession()
	if err != nil {
		return "", wUtil.NewError(err)
	}
	svc := s3.New(sess)
	avatarKey := fmt.Sprintf("%v/avatar.jpg", customerID)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(avatarBucket),
		Key:    aws.String(avatarKey),
	})
	urlStr, err := req.Presign(5 * time.Minute)
	if err != nil {
		err = wUtil.NewError(err)
	}
	return urlStr, err
}

func GetAvatar(ctx context.Context, customerID int) ([]byte, error) {
	avatarBucket := os.Getenv("aws_s3_avatar")
	if len(avatarBucket) == 0 {
		return []byte{}, wUtil.ErrorWithStr("aws_s3_avatar empty")
	}
	awsSess, err := getAWSSession()
	if err != nil {
		return []byte{}, wUtil.NewError(err)
	}
	avatarKey := fmt.Sprintf("%v/avatar.jpg", customerID)

	buf := aws.NewWriteAtBuffer([]byte{})
	downloader := s3manager.NewDownloader(awsSess)
	_, err = downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(avatarBucket),
		Key:    aws.String(avatarKey),
	})
	if err != nil {
		return []byte{}, wUtil.NewError(err)
	}
	return buf.Bytes(), nil
}

func SaveAvatar(ctx context.Context, customerID int, avatar []byte) error {
	avatarBucket := os.Getenv("aws_s3_avatar")
	if len(avatarBucket) == 0 {
		return wUtil.ErrorWithStr("aws_s3_avatar empty")
	}
	awsSess, err := getAWSSession()
	if err != nil {
		return wUtil.NewError(err)
	}
	avatarKey := fmt.Sprintf("%v/avatar.jpg", customerID)
	awsUpload := s3manager.NewUploader(awsSess)
	_, err = awsUpload.Upload(&s3manager.UploadInput{
		Bucket: aws.String(avatarBucket),
		Key:    aws.String(avatarKey),
		Body:   bytes.NewBuffer(avatar),
	})
	if err != nil {
		err = wUtil.NewError(err)
	}
	return err
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
