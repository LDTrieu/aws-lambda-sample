package awsSM

import (
	"context"
	"lambda-sample/pkg/wUtil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var (
	svcSM *secretsmanager.SecretsManager
)

func initSM(ctx context.Context) (err error) {

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
	region := os.Getenv("aws_region")
	if len(region) == 0 {
		err = wUtil.ErrorWithStr("aws_region env variable empty")
		return
	}
	cred := credentials.NewStaticCredentials(access_key, secret_access_key, "")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: cred,
	},
	)
	if err != nil {
		return err
	}

	//Create a Secrets Manager client
	svcSM = secretsmanager.New(sess,
		aws.NewConfig().WithRegion(region))
	return nil
}

func get(ctx context.Context, secretName string) (val string, err error) {
	if svcSM == nil {
		if err := initSM(ctx); err != nil {
			return "", err
		}
	}

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	// In this sample we only handle the specific exceptions for the 'GetSecretValue' API.
	// See https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html

	result, err := svcSM.GetSecretValue(input)
	if err != nil {
		return "", err
	}
	return *result.SecretString, nil
}
