package handlers

import "net/http"

func Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to ICICI"))
	w.WriteHeader(200)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	w.WriteHeader(200)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
	w.WriteHeader(200)
}
