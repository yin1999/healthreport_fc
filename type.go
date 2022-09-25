package main

import "io"

type handler interface {
	Next() (body io.ReadCloser, reqID string, err error)
	ReportError(msg string, id string)
	ReportSuccess(id string)
}
