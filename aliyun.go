//go:build aliyun || !tencent || notify

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type timerTrigger struct {
	TriggerTime string `json:"triggerTime"`
	TriggerName string `json:"triggerName"`
	Payload     string `json:"payload"`
}

const (
	fcStatus    = "X-Fc-Status"
	fcRequestId = "X-Fc-Request-Id"
	apiVersion  = "2020-11-11"
	contentType = "text/plain"
)

func regist() (handler, error) {
	address := os.Getenv("FC_RUNTIME_API")
	endpoint := fmt.Sprintf("http://%s/%s/runtime/invocation/", address, apiVersion)
	return &aliyun{
		endpoint: endpoint,
		client:   http.Client{},
	}, nil
}

type aliyun struct {
	endpoint string
	ua       string
	client   http.Client
}

func (a aliyun) Next() (body io.ReadCloser, reqID string, err error) {
	var resp *http.Response
	resp, err = http.Get(a.endpoint + "next")
	if err == nil {
		body = resp.Body
		reqID = resp.Header.Get(fcRequestId)
	}
	return
}

func (a aliyun) ReportSuccess(id string) {
	res, err := http.DefaultClient.Post(a.endpoint+id+"/response", contentType, http.NoBody)
	if err == nil {
		res.Body.Close()
	} else {
		Error.Log(err.Error() + "\n")
	}
}

func (a aliyun) ReportError(msg string, id string) {
	res, err := http.Post(a.endpoint+id+"/error",
		contentType,
		strings.NewReader(msg),
	)
	if err == nil {
		res.Body.Close()
	} else {
		Error.Log(err.Error() + "\n")
	}
}
