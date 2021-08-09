package helper

import (
	"net/http"
)

type Response interface {
	Write(w http.ResponseWriter, body []byte, status int)
}

func NewResponse(r Response) Response {
	switch r.(type) {
	case *Json:
		return &Json{}
	//you can add another response content type in here
	default:
		return &Json{}
	}
}

type Json struct{}

func (j *Json) Write(w http.ResponseWriter, body []byte, status int) {
	w.WriteHeader(status)
	w.Header().Set("content-type", "application/json")
	if _, err := w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
