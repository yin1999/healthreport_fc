package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"
	"time"

	client "github.com/yin1999/healthreport/httpclient"
)

type handler interface {
	Next() (body io.ReadCloser, reqID string, err error)
	ReportError(msg string, id string)
	ReportSuccess(id string)
}

func main() {
	h, err := regist()
	if err != nil {
		os.Exit(1)
	}
	startServe(h)
}

func init() {
	if os.Getenv("INSECURE") == "1" {
		client.SetSslVerify(false)
	}
}

func startServe(handler handler) {
	for {
		body, id, err := handler.Next()
		if err != nil {
			Error.Log("get trigger payload failed, err: %s\n", err.Error())
			continue
		}
		t := &timerTrigger{}
		err = json.NewDecoder(body).Decode(t)
		body.Close() // close body
		if err != nil {
			msg := "parse request body failed, err: " + err.Error()
			Error.Log(msg + "\n")
			handler.ReportError(msg, id)
			continue
		}
		err = punch(t.Payload)
		if err != nil {
			handler.ReportError(err.Error(), id)
		} else {
			handler.ReportSuccess(id)
		}
	}
}

func punch(payload string) error {
	account := strings.Fields(payload)
	if len(account) < 2 {
		return errors.New("invalid payload format")
	}
	err := client.Punch(context.Background(), &client.Account{Username: account[0], Password: account[1]}, 30*time.Second)
	if err != nil {
		Error.Log("account: %s punch failed, err: %s\n", account[0], err.Error())
	} else {
		Info.Log("account: %s punch success\n", account[0])
	}
	return err
}
