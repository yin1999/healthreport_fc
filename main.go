package main

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	client "github.com/yin1999/healthreport/httpclient"
)

type handler interface {
	ListenAndServe(punch func(payload string) error) error
}

func main() {
	h, err := regist()
	if err != nil {
		os.Exit(1)
	}
	if err = h.ListenAndServe(punch); err != nil {
		os.Exit(2)
	}
}

func init() {
	if os.Getenv("INSECURE") == "1" {
		client.SetSslVerify(false)
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
