package main

import (
	"context"
	"lambda-sample/lambda/sample/checker"
	lambdautil "lambda-sample/source/lambdaUtil"
	"lambda-sample/source/model"
	"lambda-sample/source/wUtil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlePost(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	api, ok := req.QueryStringParameters["api"]
	if !ok {
		return wUtil.LambdaRespError(ctx, model.CodeAppIDEmpty)

	}
	if req.HTTPMethod == http.MethodPost {
		switch api {
		case "test":
			return lambdautil.LambdaServRaw(ctx, req, checker.CheckAPI), nil
		case "login":
			return lambdautil.LambdaServRaw(ctx, req, login), nil
		}

	}
	//-----
	language := wUtil.GetCurrentLanguage(ctx, req)
	ctx = context.WithValue(ctx, model.Context.LanguageKey, language)
	//-----

	return wUtil.LambdaRespError(ctx, model.CodeAPINotExist)

}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Credentials": "true",
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "POST,HEAD,PATCH, OPTIONS, GET, PUT",
	}
	if req.HTTPMethod == http.MethodOptions {
		//log.Println("option call", req.Headers)
		return events.APIGatewayProxyResponse{
			Headers:    headers,
			StatusCode: 204,
		}, nil
	}
	resp, _ := handlePost(ctx, req)
	resp.Headers = headers
	return resp, nil
}

func main() {
	lambda.Start(handler)
}
