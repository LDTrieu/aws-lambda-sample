package model

import "lambda-sample/source/auth/sys"

type RefreshTokenReq struct {
	InstallationID  string
	RefreshJWTValue string `json:"refreshJWTValue"`
	LoginStatus     bool   `json:"loginStatus"`
}

type RefreshTokenResp struct {
	sys.SysResp
	JWTValue string `json:"jwtValue"`
}

type RefreshToken_UpdateRevokedStatusReq struct {
	InstallationID string
	Phone          string `json:"phone"`
	RevokedStatus  bool   `json:"revokedStatus"`
}

type RefreshToken_UpdateRevokedStatusResp struct {
	sys.SysResp
}
