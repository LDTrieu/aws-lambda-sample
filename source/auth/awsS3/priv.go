package awsS3

import (
	"os"

	"lambda-sample/source/wUtil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	awsSess *session.Session
)

func getAWSSession() (sess *session.Session, err error) {
	if awsSess != nil {
		return awsSess, nil
	}
	access_key := os.Getenv("aws_key")
	if len(access_key) == 0 {
		err = wUtil.ErrorWithStr("AWS_ACCESS_KEY_ID enviroment variable empty")
		return
	}
	secret_access_key := os.Getenv("aws_secret_key")
	if len(secret_access_key) == 0 {
		err = wUtil.ErrorWithStr("AWS_SECRET_ACCESS_KEY enviroment variable empty")
		return
	}
	region := os.Getenv(awsRegion)
	if len(awsRegion) == 0 {
		err = wUtil.ErrorWithStr("aws_region env variable empty")
		return
	}
	cred := credentials.NewStaticCredentials(access_key, secret_access_key, "")
	sess, err = session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: cred,
	})
	if err != nil {
		err = wUtil.NewError(err)
		return
	}
	awsSess = sess
	return
}
