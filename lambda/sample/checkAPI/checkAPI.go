package checker

import (
	"context"
	"lambda-sample/lambda/sample/util"
	"lambda-sample/pkg/model"

	"github.com/aws/aws-lambda-go/events"
)

// API for check service
func CheckAPI(ctx context.Context, apiReq events.APIGatewayProxyRequest) (interface{}, *model.FaError) {
	const mockData = `{"resultCode":"OK","data":"data_something"}`
	return util.GetClientInfo(apiReq), nil
}
