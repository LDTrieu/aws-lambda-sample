package wUtil

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"lambda-sample/source/model"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

func GetPhoneAppStore() string {
	return "0987654321"
}

func ConvertDTOPhoneNumber(phone string) string {
	phoneStr := strings.Replace(phone, "0", "+84", 1)
	return phoneStr
}

func getFileLine() (file string, line int) {
	_, file, line, _ = runtime.Caller(2)
	para := strings.Split(file, "/")
	size := len(para)
	if size > 2 {
		file = fmt.Sprintf("%v/%v", para[size-2], para[size-1])
	}
	return
}

func NewError(err error) (logErr error) {
	file, line := getFileLine()
	return fmt.Errorf(fmt.Sprintf("%v line:%v | %v", file, line, err.Error()))
}

func ErrorWithStr(message string) (logErr error) {
	file, line := getFileLine()
	return fmt.Errorf(fmt.Sprintf("%v line:%v | %v", file, line, message))
}

func StrLine(message string) string {
	file, line := getFileLine()
	return fmt.Sprintf("%v line:%v | %v", file, line, message)
}

func StrLog(a ...interface{}) string {
	file, line := getFileLine()
	return fmt.Sprint(file, " line: ", line, " | ", a)
}

func StrLogf(format string, a ...interface{}) string {
	file, line := getFileLine()
	str := fmt.Sprintf(format, a...)
	return fmt.Sprint(file, " line: ", line, " | ", str)
}

var SHA1 = func(buff []byte) (hash string) {
	hasher := sha1.New()
	hasher.Write(buff)
	hash = hex.EncodeToString(hasher.Sum(nil))
	return
}

func DateFromStr(year, month, date string) (rst time.Time, err error) {
	if len(month) == 1 {
		month = fmt.Sprintf("0%v", month)
	}
	if len(date) == 1 {
		date = fmt.Sprintf("0%v", date)
	}
	timeStr := fmt.Sprintf("%v-%v-%vT15:04:05.000Z", year, month, date)
	return time.Parse("2006-01-02T15:04:05.000Z", timeStr)
}

func LogLambda(ctx context.Context, inf interface{}) {
	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		log.Println(inf)
		return
	}
	log.Printf("%v : %v", lc.AwsRequestID, inf)
}

func GetCurrentLanguage(ctx context.Context, req events.APIGatewayProxyRequest) string {
	language := model.LanguageEN
	for k, v := range req.Headers {
		if strings.ToLower(k) == "language" {
			language = v
			if language != model.LanguageEN && language != model.LanguageVN {
				language = model.LanguageEN
			}
			break
		}
	}
	return language
}

func GetLanguageObj(ctx context.Context) *model.Language {
	lang, ok := ctx.Value(model.Context.LanguageKey).(string)
	if !ok || lang == model.LanguageVN {
		return &model.Language{
			Code: model.VNCode,
			Name: lang,
		}
	}
	return &model.Language{
		Code: model.ENCode,
		Name: lang,
	}
}

func GetLanguageByContext(ctx context.Context) string {
	language, ok := ctx.Value(model.Context.LanguageKey).(string)
	if ok {
		return language
	}
	return model.LanguageVN
}
