package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func tencentInvoke(w http.ResponseWriter, req *http.Request) {
	t := tencentTrigger{}
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		msg := "parse request body failed, err: " + err.Error()
		Error.Log(msg + "\n")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		return
	}
	err = punch(t.Payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "cannot report, err: %s", err.Error())
	}
}

type tencentTrigger struct {
	TriggerTime string `json:"Time"`
	TriggerName string `json:"TriggerName"`
	Payload     string `json:"Message"`
}
