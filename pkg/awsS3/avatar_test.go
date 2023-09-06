package awsS3

import (
	"context"
	"io/ioutil"
	"lambda-sample/pkg/wUtil"
	"net/http"
	"testing"
)

func Test_preSignAvatarURL(t *testing.T) {
	//ctx := wUtil.NewTestCtx()
	ctx := context.Background()
	const custID = 90090
	avatar, err := ioutil.ReadFile("test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	err = SaveAvatar(ctx, custID, avatar)
	if err != nil {
		t.Fatal(err)
	}
	url, err := GenPreSignAvatarURL(ctx, custID)
	if err != nil {
		t.Fatal(err)
	}
	if len(url) == 0 {
		t.Fatal("Fail gen url")
	}

	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatal("http status code invalid:", resp.StatusCode, "expect:", http.StatusOK)
	}
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if len(buff) == 0 {
		t.Fatal("file download error")
	}
	if wUtil.SHA1(avatar) != wUtil.SHA1(buff) {
		t.Fatal("Upload and download avartar ")
	}

	aBuff, err := GetAvatar(ctx, custID)
	if err != nil {
		t.Fatal(err)
	}
	if wUtil.SHA1(aBuff) != wUtil.SHA1(buff) {
		t.Fatal("Get avatar invalid")
	}
}
