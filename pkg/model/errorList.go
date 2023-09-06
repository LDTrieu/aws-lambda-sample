package model

import "fmt"

const (
	CodeJsonUnMarshal             = 2
	CodeJsonMarshal               = 3
	CodeAPINotExist               = 4
	CodeURLEmpty                  = 5
	CodeHTTPReqError              = 6
	CodeCreateTmpError            = 7
	CodeSaveTmpError              = 8
	CodeParseTimeErr              = 9
	CodeOpenTempFileErr           = 10
	CodeDBErr                     = 11

	CodeReadFileErr               = 13
	CodeBodySizeExceed            = 14
	CodeRendJWTErr                = 15
	CodeSQLNoRow                  = 16
	CodeAWSLoadErr                = 17
	CodeAvatarErr                 = 18
	CodeHVActionNotPass           = 19
	CodeNoInstallID               = 20
	CodeSessionExpired            = 21
	CodeInvalidQueryParam         = 22
	CodeFiledEmpty                = 23
	CodeUserNameEsignEmpty        = 24
	CodePwdEsignEmpty             = 25
	CodeChannelPartnerEmpty       = 26
	CodeDataNotExist              = 27
	CodeSendManualReviewErr       = 28
	CodeMultiDevice               = 29
	CodeDateStrInvalid            = 30
	CodeCreateHttpReqErr          = 31

	CodeRefreshTokenErr           = 33
	CodeSendHttpReqErr            = 34
	CodeLinkOldToNewTempFileErr   = 35
	CodeSaveS3Err                 = 36
	CodeRefreshJWTExpire          = 37
	CodeVerifyJWTErr              = 38
	CodeRequestOTPToActiveCardErr = 39
	CodeJWTExpire                 = 40
	CodeAESErr                    = 41
	CodeRefreshJWTRevoked         = 42
	CodeAuthenEmpty               = 43
	CodeInvalidFormat             = 44

	//firebase error
	CodeFbNoDocument = 50
	CodeFbError      = 51

	CodeAppKeyEmpty             = 300
	CodeAppSecretEmpty          = 301
	CodeAppIDEmpty              = 302
	CodeCheckNIDErr             = 303
	CodeCheckCustomerProfileErr = 304

	ParameterInvalid = 1002
)

var (
	ErrAPINotExist MessErr = MessErr{
		Vn: fmt.Errorf("api không tồn tại"),
		En: fmt.Errorf("api not exist"),
	}

	//aws
	ErrS3BucketEmpty   = fmt.Errorf("aws_s3_bucket env variable empty")
	ErrJWTExpire       = fmt.Errorf("JWT expired")
	ErrAuthHeaderEmpty = fmt.Errorf("auth header empty")
)

type MessErr struct {
	Vn error
	En error
}

func (it *MessErr) GetLangErr(langName string) error {
	if langName == LanguageVN {
		return it.Vn
	}

	return it.En
}
