package main

import (
	"log"
	"net/http"
	"time"
)

func accessLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("Falcon-Push: %s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))
	})
}
