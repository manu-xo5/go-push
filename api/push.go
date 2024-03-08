package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"firebase.google.com/go/messaging"
	"github.com/manu-xo5/go-push/fire"
)

type pushHandler struct{}

func (p *pushHandler) Handler(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return p.get(w, r)
	}

	return fmt.Errorf("method not allowed")
}

func (p *pushHandler) get(w http.ResponseWriter, r *http.Request) error {
	println(r.RequestURI + " sending to emeox electron")

	app, err := fire.GetApp()
	if err != nil {
		return fmt.Errorf("firebase service unavaliable")
	}

	fcm, err := app.Messaging(context.Background())
	if err != nil {
		return fmt.Errorf("firebase service unavaliable")
	}

	res, err := fcm.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: "GO FCM",
			Body:  `{ "method": "log", "params": ["hello world", "rpc"] }`,
		},
		Token: "do9z1D2AJHc:APA91bEWwxBdHw-LF9zKXUvYQqjixhlGmfLoob5UWdv7b0aqLg9GC-jUIrwqFhj0O3h1gY1F8qUBtXtiQlesd6IUuV1hfQxQq44mbU-jlbrdRCIL6U2DQdJYMmclyFmz-OqGOwYxXNRs",
	})

	if err != nil {
		return fmt.Errorf("firebase failed to send push")
	}

	json.NewEncoder(w).Encode(res)

	return nil
}

var PushHandle = pushHandler{}
