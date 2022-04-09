package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type aliyunTrigger struct {
	TriggerTime string `json:"triggerTime"`
	TriggerName string `json:"triggerName"`
	Payload     string `json:"payload"`
}

const fcStatus = "X-Fc-Status"

func aliyunInvoke(w http.ResponseWriter, req *http.Request) {
	t := aliyunTrigger{}
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		msg := "parse request body failed, err: " + err.Error()
		Error.Log(msg + "\n")
		w.Header().Add(fcStatus, "404")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		return
	}
	err = punch(t.Payload)
	if err != nil {
		w.Header().Add(fcStatus, "404")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "cannot report, err: %s", err.Error())
	}
}
