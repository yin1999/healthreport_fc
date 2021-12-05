//go:build aliyun || !tencent

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type timerTrigger struct {
	TriggerTime string `json:"triggerTime"`
	TriggerName string `json:"triggerName"`
	Payload     string `json:"payload"`
}

const (
	fcStatus = "X-Fc-Status"
)

func init() {
	h = &aliyun{}
}

type aliyun struct {
	port string
}

func (a *aliyun) Init() error {
	a.port = os.Getenv("FC_SERVER_PORT")
	if a.port == "" {
		a.port = "9000"
	}
	return nil
}

func (a aliyun) ListenAndServe(punch func(payload string) error) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get("X-Fc-Request-Id")
		fmt.Printf("FC Invoke Start RequestId: %s\n", requestId)
		defer fmt.Printf("FC Invoke End RequestId: %s\n", requestId)

		t := &timerTrigger{}
		err := json.NewDecoder(r.Body).Decode(t)

		if err != nil {
			Fatal.Log(requestId, "error payload format\n")
			w.Header().Set(fcStatus, "404")
			w.Write([]byte("error payload format"))
			return
		}
		err = punch(t.Payload)
		if err != nil {
			Error.Log(requestId, "Punch Failed: %s\n", err.Error())
			w.Header().Set(fcStatus, "404")
			fmt.Fprintf(w, "Punch Failed: %s", err.Error())
		} else {
			Info.Log(requestId, "Punch Success\n")
			w.Write([]byte("success"))
		}
	})
	return http.ListenAndServe(":"+a.port, nil)
}
