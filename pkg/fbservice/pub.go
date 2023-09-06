package fbservice

import (
	"context"
	"lambda-sample/pkg/model"
	"lambda-sample/pkg/wUtil"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

func SubscribeTopic(ctx context.Context, topic, fbToken string) error {
	f := func(app *firebase.App) error {
		client, err := app.Messaging(ctx)
		if err != nil {
			return wUtil.NewError(err)
		}
		resp, err := client.SubscribeToTopic(ctx, []string{fbToken}, topic)
		if err != nil {
			return wUtil.NewError(err)
		}
		if resp.SuccessCount == 0 {
			return wUtil.ErrorWithStr("subscribe topic fail")
		}
		return nil
	}
	return execute(f)
}

func UnsubscribeTopic(ctx context.Context, topic, fbToken string) error {
	f := func(app *firebase.App) error {
		client, err := app.Messaging(ctx)
		if err != nil {
			return wUtil.NewError(err)
		}
		_, err = client.UnsubscribeFromTopic(ctx, []string{fbToken}, topic)
		if err != nil {
			err = wUtil.NewError(err)
		}
		return err
	}
	return execute(f)
}

func NotifyTopic(ctx context.Context, topic, title, body string) error {
	f := func(app *firebase.App) error {
		client, err := app.Messaging(ctx)
		if err != nil {
			return wUtil.NewError(err)
		}
		message := &messaging.Message{
			Notification: &messaging.Notification{
				Title: title,
				Body:  body,
			},
			Topic: topic,
		}
		_, err = client.Send(ctx, message)
		if err != nil {
			err = wUtil.NewError(err)
		}
		return err
	}
	return execute(f)
}

func RunFS(ctx context.Context, f func(client *firestore.Client) (err *model.FaError)) *model.FaError {
	if fsClient == nil {
		initFbApp(context.Background())
	}
	err := f(fsClient)
	return err
}
