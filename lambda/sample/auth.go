package main

import (
	"context"
	"encoding/json"
	"fmt"
	"lambda-sample/pkg/wUtil"

	"lambda-sample/pkg/model"

	"github.com/aws/aws-lambda-go/events"
)

func login(ctx context.Context, apiReq events.APIGatewayProxyRequest) (
	resp interface{}, faErr *model.FaError) {
	req := &struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	lang := wUtil.GetLanguageObj(ctx)
	if err := json.Unmarshal([]byte(apiReq.Body), req); err != nil {
		//wlog.LogSystem(ctx, "login", wUtil.StrLogf("json mashar error: %v", err))
		return nil, &model.FaError{
			Code:     model.CodeJsonUnMarshal,
			Message:  err.Error(),
			Err:      wUtil.NewError(err),
			LangCode: lang.Name,
		}
	}
	if len(req.Username) <= 0 || len(req.Username) >= 1024 ||
		len(req.Password) <= 0 || len(req.Password) >= 1024 {
		return nil, &model.FaError{
			Code:    model.ParameterInvalid,
			Message: "Prameter invalid",
			Err:     fmt.Errorf("parameter invalid"),
		}
	}
	//resp, faErr = permission.Login(ctx, req.Username, req.Password)
	return
}
