package fbservice

import (
	"context"
	"lambda-sample/source/model"
	"lambda-sample/source/wUtil"

	firebase "firebase.google.com/go"
)

type notify struct{}

var Notify *notify

func (me *notify) SendTopic(ctx context.Context, req *model.CustomerNotify) error {
	return execute(func(app *firebase.App) error {
		client, err := app.Messaging(ctx)
		if err != nil {
			return wUtil.NewError(err)
		}
		message := newNotifyMess(req.Title, req.ShortText, req.ImgURL, req.VideoURL)
		message.Topic = req.NotifyType
		_, err = client.Send(ctx, message)
		if err != nil {
			err = wUtil.NewError(err)
		}
		return err
	})
}
