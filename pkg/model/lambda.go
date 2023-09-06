package model

import "fmt"

type LambdaResp struct {
	ReqID  string      `json:"reqID"`
	Result interface{} `json:"result"`
}

type LambdaErr struct {
	ReqID  string   `json:"reqID"`
	Result *FaError `json:"result"`
}

func (ins *LambdaErr) Error() string {
	return fmt.Sprintf("ReqID:%v Code:%v", ins.ReqID, ins.Result.Code)
}
