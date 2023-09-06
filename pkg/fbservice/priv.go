package fbservice

import (
	"context"
	"lambda-sample/pkg/auth/awsS3"
	"lambda-sample/pkg/sercfg"
	"lambda-sample/pkg/wUtil"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var (
	fbApp    *firebase.App
	fsClient *firestore.Client
)

func initFbApp(ctx context.Context) (*firebase.App, error) {
	fbKey, err := awsS3.LoadKSFile(context.Background(), "fb_key.json")
	if err != nil {
		return nil, err
	}
	// log.Println("fbKey:", len(fbKey))
	projectID := sercfg.Get(ctx, "fb_project_id")
	if len(projectID) == 0 {
		//dev env
		projectID = "lamda-sample-96e11"
	}
	dbURL := sercfg.Get(ctx, "fb_rtdb_url")
	if len(dbURL) == 0 {
		//dev env
		dbURL = "https://lamdasample.firebaseio.com/"
	}
	// log.Println(wUtil.StrLog("projectID:", projectID, "dbURL:", dbURL))

	// home := os.Getenv("HOME")
	// file := fmt.Sprintf("%v/Downloads/fb_key(1).json", home)
	// fbKey, err := ioutil.ReadFile(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	conf := &firebase.Config{
		DatabaseURL: dbURL,
		ProjectID:   projectID,
	}
	opt := option.WithCredentialsJSON(fbKey)
	// t0 := time.Now()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		// log.Println("int fb err:", err)
		return nil, wUtil.NewError(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, wUtil.NewError(err)
	}
	fsClient = client
	// log.Println("init fb ok:", time.Until(t0).Seconds())
	return app, nil
}

func execute(f func(app *firebase.App) error) (err error) {
	if fbApp == nil {
		// log.Println("init firebase")
		fbApp, err = initFbApp(context.Background())
		if err != nil {
			return wUtil.NewError(err)
		}
	}
	return f(fbApp)
}

func newNotifyMess(title, body, imgUrl, videUrl string) (mess *messaging.Message) {
	mess = &messaging.Message{
		Notification: &messaging.Notification{
			Title:    title,
			Body:     body,
			ImageURL: imgUrl,
		},
		Android: &messaging.AndroidConfig{
			Priority:     "high",
			Notification: &messaging.AndroidNotification{},
		},
	}
	if len(videUrl) > 0 {
		mess.Data = make(map[string]string)
		mess.Data["video"] = videUrl
	}
	return
}
