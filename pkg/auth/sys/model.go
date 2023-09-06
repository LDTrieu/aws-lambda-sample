package sys

type SysResp struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	LangCode string `json:"langCode"`
}
