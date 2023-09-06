package sercfg

import (
	"context"
	"encoding/json"
	"lambda-sample/pkg/auth/awsS3"

	"log"
	"os"
)

var serCfg map[string]string
var cfgVersion string

func Get(ctx context.Context, key string) string {
	cVer := os.Getenv("cfgVersion")
	if len(serCfg) == 0 || cfgVersion != cVer {
		refesh(ctx)
		cfgVersion = cVer
	}
	return serCfg[key]
}

func refesh(ctx context.Context) {
	keyFile := os.Getenv("cfgfile")
	if len(keyFile) == 0 {
		log.Fatal("key file empty")
		return
	}
	buff, err := awsS3.LoadKSFile(ctx, keyFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err = json.Unmarshal(buff, &serCfg); err != nil {
		log.Fatal(err)
	}
}
