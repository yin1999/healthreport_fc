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
	Init() error
	ListenAndServe(punch func(payload string) error) error
}

var h handler

func main() {
	err := h.Init()
	if err != nil {
		os.Exit(1)
	}
	if err = h.ListenAndServe(punch); err != nil {
		os.Exit(2)
	}
}

func punch(payload string) error {
	account := strings.Fields(payload)
	if len(account) < 2 {
		return errors.New("invalid payload format")
	}
	return client.Punch(context.Background(), &client.Account{Username: account[0], Password: account[1]}, 30*time.Second)
}
