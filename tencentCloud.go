//go:build tencent

package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

const contentType = "text/plain"

func regist() (handler, error) {
	host := os.Getenv("SCF_RUNTIME_API")
	port := os.Getenv("SCF_RUNTIME_API_PORT")
	reportAPI := "http://" + host + ":" + port
	res, err := http.Post(reportAPI+"/runtime/init/ready", contentType, http.NoBody)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	return &tencent{
		reportAPI: reportAPI,
	}, err
}

type tencent struct {
	reportAPI string
}

type timerTrigger struct {
	TriggerTime string `json:"Time"`
	TriggerName string `json:"TriggerName"`
	Payload     string `json:"Message"`
}

func (h tencent) ListenAndServe(punch func(payload string) error) error {
	for {
		res, err := http.Get(h.reportAPI + "/runtime/invocation/next")
		if err != nil {
			Error.Log("get trigger payload failed, err: %s\n", err.Error())
			h.reportError("", "get payload failed, err: "+err.Error())
		}
		requestId := res.Header.Get("Request_id")
		dec := json.NewDecoder(res.Body)
		t := &timerTrigger{}
		err = dec.Decode(t)
		res.Body.Close() // close body
		if err != nil {
			msg := "parse request body failed, err: " + err.Error()
			Error.Log(requestId, msg+"\n")
			h.reportError(requestId, msg)
		}
		err = punch(t.Payload)
		if err != nil {
			msg := "punch failed, err: " + err.Error()
			Error.Log(requestId, msg+"\n")
			h.reportError(requestId, msg)
		} else {
			Info.Log(requestId, "punch success\n")
			res, err := http.DefaultClient.Post(h.reportAPI+"/runtime/invocation/response", contentType, strings.NewReader(requestId))
			if err == nil {
				res.Body.Close()
			} else {
				Fatal.Log(requestId, err.Error()+"\n")
			}
		}
	}
}

func (h tencent) reportError(requestId, msg string) {
	res, err := http.DefaultClient.Post(h.reportAPI+"/runtime/invocation/error",
		contentType,
		strings.NewReader(msg),
	)
	if err == nil {
		res.Body.Close()
	} else {
		Error.Log(requestId, err.Error()+"\n")
	}
}
