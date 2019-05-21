package main

import (
	"log"
	"net/http"
	"strings"
)

func Router(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path, r.RemoteAddr)
	if r.Method == "OPTIONS" {
		writeAccessOptions(w)
		return
	}
	path := strings.TrimPrefix(r.URL.Path, "/api/")
	parts := strings.Split(path, "/")
	if len(parts) < 1 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	cmd := parts[0]
	switch cmd {
	case "hello":
		ServeJSON("Hello World 2!", w)
	}
}