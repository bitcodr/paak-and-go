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
	//you can add another response content type in here and implement it
	default:
		return &Json{}
	}
}

type Json struct{}

func (j *Json) Write(w http.ResponseWriter, body []byte, status int) {
	//todo implement better structure for errors
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
