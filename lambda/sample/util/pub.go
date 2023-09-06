package util

import (
	//lambdaModel "lambda-sample/lambda/sample/model"
	lambdaModel "lambda-sample/lambda/sample/model"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func GetClientInfo(reqApi events.APIGatewayProxyRequest) (info *lambdaModel.ClientInfo) {
	info = &lambdaModel.ClientInfo{}
	para := strings.Split(reqApi.Headers["X-Forwarded-For"], ",")
	info.Ip = []string{strings.TrimSpace(para[0]),
		strings.TrimSpace(para[1])}

	info.Country = reqApi.Headers["CloudFront-Viewer-Country"]
	info.UserAgent = reqApi.Headers["User-Agent"]
	for k, v := range reqApi.Headers {
		if !strings.Contains(k, "Viewer") {
			continue
		}
		if v == "true" {
			info.CLientType = k
			break
		}
	}
	return
}
