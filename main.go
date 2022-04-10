package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	client "github.com/yin1999/healthreport/httpclient"
	"github.com/yin1999/healthreport/utils/captcha"
)

type handler interface {
	Next() (body io.ReadCloser, reqID string, err error)
	ReportError(msg string, id string)
	ReportSuccess(id string)
}

func main() {
	captcha.Init()
	defer captcha.Close()
	http.HandleFunc("/invoke", aliyunInvoke)
	http.HandleFunc("/event-invoke", tencentInvoke)
	http.ListenAndServe("0.0.0.0:9000", nil)
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
