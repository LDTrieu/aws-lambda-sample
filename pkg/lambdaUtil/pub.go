package lambdautil

import (
	"context"
	"encoding/json"
	"fmt"
	"lambda-sample/pkg/auth/sys"
	"lambda-sample/pkg/model"
	"lambda-sample/pkg/wUtil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LambdaServRaw(ctx context.Context, req events.APIGatewayProxyRequest,
	act func(context.Context, events.APIGatewayProxyRequest) (interface{}, *model.FaError)) events.APIGatewayProxyResponse {

	var reqID string
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		reqID = lc.AwsRequestID
	} else {
		reqID = primitive.NewObjectID().Hex()
	}
	rst, faErr := act(ctx, req)
	if faErr != nil {
		if faErr.Err != nil {
			//wlog.LogSystem(ctx, "LambdaError", faErr.Error())
		}
		lResp, _ := LambdaRespError(ctx, faErr.Code, faErr.Message)
		return lResp
	}
	//---
	resp := &struct {
		sys.SysResp
		ReqID  string      `json:"reqID"`
		Result interface{} `json:"result,omitempty"`
	}{
		ReqID:  reqID,
		Result: rst,
	}
	resp.LangCode = wUtil.GetLanguageByContext(ctx)
	buff, err := json.Marshal(resp)
	if err != nil {
		//wlog.LogSystem(ctx, "LambdaErr", wUtil.NewError(err).Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(buff),
	}
}

func LambdaRespError(ctx context.Context, code int,
	message string) (events.APIGatewayProxyResponse, error) {
	resp := &struct {
		sys.SysResp
		ReqID string `json:"reqId"`
	}{}
	resp.Code = code
	resp.Message = message
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		resp.ReqID = lc.AwsRequestID
	} else {
		resp.ReqID = primitive.NewObjectID().Hex()
	}

	body, err := json.Marshal(resp)
	if err != nil {
		//wlog.LogSystem(ctx, "LambdaErr", err.Error())
		bodyStr := fmt.Sprintf("{\"reqID\":\"%v\",\"code:\"%v\"}",
			resp.ReqID, resp.Code)
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
