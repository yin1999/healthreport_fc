//go:build notify

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/yin1999/healthreport/utils/email"
)

var emailCfg *email.Config

func init() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	emailCfg = &email.Config{
		To: strings.Split(os.Getenv("TO"), ","),
		SMTP: email.SmtpConfig{
			Host:     os.Getenv("HOST"),
			Port:     port,
			TLS:      os.Getenv("TLS") == "true",
			Username: os.Getenv("USERNAME"),
			Password: os.Getenv("PASSWORD"),
		},
	}
}

func main() {
	h, err := regist()
	if err != nil {
		os.Exit(1)
	}
	startServe(h)
}

type message struct {
	ResponsePayload string
}

func startServe(handler handler) {
	for {
		body, id, err := handler.Next()
		if err != nil {
			Error.Log("get trigger payload failed, err: %s\n", err.Error())
			continue
		}
		t := &message{}
		err = json.NewDecoder(body).Decode(t)
		body.Close() // close body
		if err != nil {
			msg := "parse request body failed, err: " + err.Error()
			Error.Log(msg + "\n")
			handler.ReportError(msg, id)
			continue
		}
		err = emailCfg.Send("打卡状态推送", fmt.Sprintf("打卡状态推送-%s",
			time.Now().Format("2006-01-02")),
			fmt.Sprintf("打卡失败: %s ", t.ResponsePayload))
		if err != nil {
			handler.ReportError(err.Error(), id)
		} else {
			handler.ReportSuccess(id)
		}
	}
}
