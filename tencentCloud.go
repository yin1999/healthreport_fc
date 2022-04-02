//go:build tencent

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const contentType = "text/plain"

func regist() (handler, error) {
	host := os.Getenv("SCF_RUNTIME_API")
	port := os.Getenv("SCF_RUNTIME_API_PORT")
	endpoint := fmt.Sprintf("http://%s:%s/", host, port)
	res, err := http.Post(endpoint+"runtime/init/ready", contentType, http.NoBody)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	return &tencent{
		endpoint: endpoint,
	}, err
}

type tencent struct {
	endpoint string
}

type timerTrigger struct {
	TriggerTime string `json:"Time"`
	TriggerName string `json:"TriggerName"`
	Payload     string `json:"Message"`
}

func (h tencent) Next() (body io.ReadCloser, id string, err error) {
	var resp *http.Response
	resp, err = http.Get(h.endpoint + "runtime/invocation/next")
	if err == nil {
		body = resp.Body
		id = resp.Header.Get("Request_id")
	}
	return
}

func (h tencent) ReportSuccess(id string) {
	res, err := http.DefaultClient.Post(h.endpoint+"runtime/invocation/response", contentType, strings.NewReader(id))
	if err == nil {
		res.Body.Close()
	} else {
		Fatal.Log(err.Error() + "\n")
	}
}

func (h tencent) ReportError(msg string, _ string) {
	res, err := http.DefaultClient.Post(h.endpoint+"runtime/invocation/error",
		contentType,
		strings.NewReader(msg),
	)
	if err == nil {
		res.Body.Close()
	} else {
		Error.Log(err.Error() + "\n")
	}
}
