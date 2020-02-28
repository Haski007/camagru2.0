package logger

import (
	"log"
	"net/http"
)

// JSON writes to logs method of response and ip of user and set json header.
func JSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		log.Printf("server [net/http] method %s connection from [%v]", r.Method, r.RemoteAddr)

		next.ServeHTTP(w, r)
	}
}
