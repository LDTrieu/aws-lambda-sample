package auth

import (
	"context"
	"crypto/rsa"
	"lambda-sample/pkg/auth/awsS3"
	"lambda-sample/pkg/sercfg"
	"lambda-sample/pkg/wUtil"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dgrijalva/jwt-go"
)

const JWTHeader = " apijwt"

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

type SampleClaim struct {
	*jwt.StandardClaims
	TokenType  string
	CustomerID int
	Refesh     bool
}

func initJWTKey() (err error) {
	ctx := context.Background()
	privKeyName := sercfg.Get(ctx, "jwtPriv")
	privBuff, err := awsS3.LoadKSFile(ctx, privKeyName)
	if err != nil {
		return err
	}
	pubKeyName := sercfg.Get(ctx, "jwtPub")
	pubBuff, err := awsS3.LoadKSFile(ctx, pubKeyName)
	if err != nil {
		return err
	}

	if signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privBuff); err != nil {
		return wUtil.NewError(err)
	}
	if verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(pubBuff); err != nil {
		return wUtil.NewError(err)
	}
	return
}

func createJWT(customerID int, isRefesh bool) (string, error) {

	if signKey == nil {
		if err := initJWTKey(); err != nil {
			return "", wUtil.NewError(err)
		}
	}
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims = &SampleClaim{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
		"level1",
		customerID,
		isRefesh,
	}
	jwtStr, err := token.SignedString(signKey)
	if err != nil {
		err = wUtil.NewError(err)
	}
	return jwtStr, err
}

var AuthReq = func(request interface{}) (err error) {
	req, ok := request.(events.APIGatewayProxyRequest)
	if !ok {
		err = wUtil.ErrorWithStr("Invalid data struct")
		return
	}
	jwtVal, ok := req.Headers[JWTHeader]
	if !ok {
		err = wUtil.ErrorWithStr("JWT not found")
		return
	}
	if len(jwtVal) < 10 {
		err = wUtil.ErrorWithStr("JWT invalid")
	}
	return
}

var CreatJWT = func(customerID int) (string, error) {
	return createJWT(customerID, false)
}

var CreateJWTRefesh = func(customerID int) (string, error) {
	return createJWT(customerID, true)
}

var VerifyJWT = func(jwtStr string) (int, error) {
	token, err := jwt.ParseWithClaims(jwtStr, &SampleClaim{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		return -1, wUtil.NewError(err)
	}
	return token.Claims.(*SampleClaim).CustomerID, nil
}

var VerifyRefeshJWT = func(jwtStr string) (int, error) {
	token, err := jwt.ParseWithClaims(jwtStr, &SampleClaim{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		return -1, wUtil.NewError(err)
	}

	SampleClaim, ok := token.Claims.(*SampleClaim)
	if !ok {
		return -1, wUtil.ErrorWithStr("Invalid claim format")
	}
	if !SampleClaim.Refesh {
		return -1, wUtil.ErrorWithStr("This isn't refesh token")
	}
	return SampleClaim.CustomerID, nil
}
