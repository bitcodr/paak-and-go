package helper

import (
	"net/http"
)

//Response interface to create multiple type of response like json, msgpack, etc
type Response interface {
	Write(w http.ResponseWriter, body []byte, status int)
}

//NewResponse - init response model
func NewResponse(r Response) Response {
	switch r.(type) {
	case *Json:
		return &Json{}
	//you can add another response content type in here and implement it
	default:
		return &Json{}
	}
}

//Json - use it if you want to pass json content type to the response
type Json struct{}

//Write to http response body with json content type
func (j *Json) Write(w http.ResponseWriter, body []byte, status int) {
	//todo implement better structure for errors
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
