package middleware

import (
	"log"
	"net/http"
)

//Logging we add metrics and tracing in here for monitoring purpose in here
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//add tracer and metrics
		log.Println(r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
