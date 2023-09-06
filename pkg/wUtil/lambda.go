package wUtil

import (
	"context"
	"encoding/json"
	"fmt"
	"lambda-sample/pkg/model"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LambdaServJSON(ctx context.Context, req events.APIGatewayProxyRequest, act func(context.Context, []byte) (rst interface{}, err error)) events.APIGatewayProxyResponse {
	rst, err := act(ctx, []byte(req.Body))
	if err != nil {
		LogLambda(ctx, err)
	}
	if rst == nil {
		messErr := "result data is nil"
		if err != nil {
			messErr = err.Error()
		}
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       messErr,
		}
	}
	var reqID string
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		reqID = lc.AwsRequestID
	} else {
		reqID = primitive.NewObjectID().Hex()
	}
	resp := &struct {
		ReqID  string      `json:"reqID"`
		Result interface{} `json:"result"`
	}{
		ReqID:  reqID,
		Result: rst,
	}
	buff, err := json.Marshal(resp)
	if err != nil {
		LogLambda(ctx, err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(buff),
	}
}

func LambdaError(ctx context.Context, code int, err error) (events.APIGatewayProxyResponse, error) {
	LogLambda(ctx, err)
	return LambdaRespError(ctx, code)
}

func LambdaRespError(ctx context.Context, code int) (events.APIGatewayProxyResponse, error) {
	resp := &model.LambdaErr{}
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		resp.ReqID = lc.AwsRequestID
	} else {
		resp.ReqID = primitive.NewObjectID().Hex()
	}
	//---
	resp.Result = &model.FaError{
		Code:    code,
		Message: getMessErr(ctx, code),
	}

	body, err := json.Marshal(resp)
	if err != nil {
		LogLambda(ctx, NewError(err))
		bodyStr := fmt.Sprintf("{\"reqID\":\"%v\",\"code:\"%v\"}",
			resp.ReqID, resp.Result.Code)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       bodyStr,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       string(body),
	}, nil
}

func getMessErr(ctx context.Context, code int) string {
	language := GetLanguageByContext(ctx)
	if code == model.CodeAPINotExist {
		return model.ErrAPINotExist.GetLangErr(language).Error()
	}
	return ""
}
