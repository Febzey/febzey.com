package controllers

import (
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("Hello"))
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("Test"))
}
