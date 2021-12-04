package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	client "github.com/yin1999/healthreport/httpclient"
)

type timerTrigger struct {
	TriggerTime string `json:"triggerTime"`
	TriggerName string `json:"triggerName"`
	Payload     string `json:"payload"`
}

const (
	fcStatus = "x-fc-status"
)

func main() {
	port := os.Getenv("FC_SERVER_PORT")
	if port == "" {
		port = "9000"
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("x-fc-request-id")
	fmt.Printf("FC Invoke Start RequestId: %s\n", requestID)
	defer fmt.Printf("FC Invoke End RequestId: %s\n", requestID)

	var t = &timerTrigger{}
	err := json.NewDecoder(r.Body).Decode(t)

	if err != nil {
		Log(Fatal, requestID, "error payload format\n")
		w.Header().Set(fcStatus, "404")
		w.Write([]byte("error payload format"))
		return
	}
	account := strings.Fields(t.Payload)
	if len(account) < 2 {
		Log(Fatal, requestID, "error account data\n")
		w.Header().Set(fcStatus, "404")
		w.Write([]byte("error account data"))
		return
	}
	err = client.Punch(context.Background(), &client.Account{Username: account[0], Password: account[1]}, 30*time.Second)
	if err != nil {
		Log(Error, requestID, "Punch Failed: %s\n", err.Error())
		w.Header().Set(fcStatus, "404")
		fmt.Fprintf(w, "Punch Failed: %s", err.Error())
	} else {
		Log(Info, requestID, "Punch Success\n")
		w.Write([]byte("success"))
	}
}

// Level log level
type Level int

const (
	None Level = iota
	Info
	Error
	Fatal
)

var level2string = [...]string{"", "Info", "Error", "Fatal"}

func (l Level) String() string {
	return level2string[l]
}

func Log(level Level, requestID, format string, v ...interface{}) {
	t := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	prefix := fmt.Sprintf("%s RequestID: %s ", t, requestID)
	if level != None {
		prefix += fmt.Sprintf("[%s] ", level)
	}
	fmt.Fprintf(os.Stderr, prefix+format, v...)
}
