package auth

import (
	"lambda-sample/source/model"
	"time"
)

const (
	AuthHeader = "auth"
)

type BASEGuestClaim struct {
	InstallationID string                 `json:"installationID"`
	CustomerID     int64                  `json:"customerID"`
	IssueAt        int64                  `json:"iat"`
	ExpireAt       int64                  `json:"exp"`
	Issuer         string                 `json:"issuer"`
	IsRefeshJwt    bool                   `json:"isRefeshJwt"`
	Act            map[string]interface{} `json:"act"`
}

type JWTResult struct {
	Code          int    `json:"code"`
	JWTVal        string `json:"jwt"`
	RefreshJWTVal string `json:"refreshJwt"`
}

func (ins *BASEGuestClaim) Valid() error {
	if time.Now().Unix()-ins.ExpireAt > 0 {
		return model.ErrJWTExpire
	}
	return nil
}
