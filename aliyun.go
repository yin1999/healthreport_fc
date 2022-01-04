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

func regist() (handler, error) {
	port := os.Getenv("FC_SERVER_PORT")
	if port == "" {
		port = "9000"
	}
	return &aliyun{
		port: port,
	}, nil
}

type aliyun struct {
	port string
}

func (a aliyun) ListenAndServe(punch func(payload string) error) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get("X-Fc-Request-Id")
		fmt.Printf("FC Invoke Start RequestId: %s\n", requestId)
		defer fmt.Printf("FC Invoke End RequestId: %s\n", requestId)

		t := &timerTrigger{}
		err := json.NewDecoder(r.Body).Decode(t)

		if err != nil {
			Fatal.Log("error payload format\n")
			w.Header().Set(fcStatus, "404")
			w.Write([]byte("error payload format"))
			return
		}
		err = punch(t.Payload)
		if err != nil {
			w.Header().Set(fcStatus, "404")
			fmt.Fprintf(w, "Punch Failed: %s", err.Error())
		} else {
			w.Write([]byte("success"))
		}
	})
	return http.ListenAndServe(":"+a.port, nil)
}
