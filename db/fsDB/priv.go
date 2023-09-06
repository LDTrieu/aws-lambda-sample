package fsdb

import (
	"context"
	"fmt"
	"lambda-sample/pkg/fbservice"
	"lambda-sample/pkg/model"
	"lambda-sample/pkg/wUtil"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Set `obj` for collection `collName`, with document id = `id` and tracking request by installation id `installID`
func set(ctx context.Context, installID, id, collName string, obj interface{}) *model.FaError {
	var (
		t0         = time.Now()
		writeCount int
		action     = fmt.Sprintf("set for collection %s", collName)
	)
	wErr := fbservice.RunFS(ctx, func(client *firestore.Client) (err *model.FaError) {
		coll := client.Collection(collName)
		_, errFb := coll.Doc(id).Set(ctx, obj)
		if errFb != nil {
			return &model.FaError{
				Code:     model.CodeFbError,
				Message:  "error query",
				LangCode: model.LanguageEN,
				Err:      err,
			}
		}
		writeCount = 1
		return nil
	})
	dur := time.Since(t0).Milliseconds()
	logFb(ctx, installID, action, writeCount, dur)
	return wErr
}

// Add `obj` for collection `collName` and tracking request by installation id `installID`
//
// Response `id` document id in collection, and `wErr` error
func add(ctx context.Context, installID, coll string, obj interface{}) (id string, wErr *model.FaError) {
	if obj == nil {
		return "", model.NewFaErr(model.ParameterInvalid, "obj parameter is nil",
			fmt.Errorf("obj parameter is nil"))
	}
	var (
		t0         = time.Now()
		writeCount int
		action     = fmt.Sprintf("add for collection %s", coll)
	)
	wErr = fbservice.RunFS(ctx, func(client *firestore.Client) *model.FaError {
		coll := client.Collection(coll)
		ref, _, err := coll.Add(ctx, obj)
		if err != nil {
			return &model.FaError{
				Code:     model.CodeFbError,
				Message:  "error query",
				LangCode: model.LanguageEN,
				Err:      err,
			}
		}
		writeCount = 1
		id = ref.ID
		return nil
	})
	dur := time.Since(t0).Milliseconds()
	logFb(ctx, installID, action, writeCount, dur)
	return
}

// Get document at collection `coll` with document id = `id` and tracking request by installation id `installID`
func getByID(ctx context.Context, obj interface{}, installID, coll, id string) (
	wErr *model.FaError) {
	if obj == nil {
		return model.NewFaErr(model.ParameterInvalid, "obj parameter is nil",
			fmt.Errorf("obj parameter is nil"))
	}
	var (
		t0        = time.Now()
		action    = fmt.Sprintf("get document %s for collection %s", id, coll)
		readCount = 0
	)
	wErr = fbservice.RunFS(ctx, func(client *firestore.Client) *model.FaError {
		docRef, err := client.Collection(coll).Doc(id).Get(ctx)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return &model.FaError{
					Code:     model.CodeFbNoDocument,
					Message:  "some requested entity was not found",
					LangCode: model.LanguageEN,
					Err:      err,
				}
			}
			return &model.FaError{
				Code:     model.CodeFbError,
				Message:  "error query",
				LangCode: model.LanguageEN,
				Err:      err,
			}
		}
		if err = docRef.DataTo(obj); err != nil {
			return &model.FaError{
				Code:     model.CodeJsonUnMarshal,
				Message:  "json unmarshal failed",
				LangCode: model.LanguageEN,
				Err:      err,
			}
		}
		readCount = 1
		return nil
	})
	dur := time.Since(t0).Milliseconds()
	logFb(ctx, installID, action, readCount, dur)
	return
}

// Write Log System for `installID` with `action` and `count` count of read or write in duration time `dur`
func logFb(ctx context.Context, installID, action string, count int, dur int64) {
	message := fmt.Sprintf("%s with count tracking: %d in duration time %d milliseconds",
		wUtil.StrLine("set"), count, dur)
	//wlog.LogSystem(ctx, installID, message)
	log.Println(message)
}
