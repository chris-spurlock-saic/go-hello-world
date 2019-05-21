package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	args := GetArgs()
	h := http.NewServeMux()
	h.HandleFunc("/api/", Router)
	log.Println("Starting server at port:", args.Port)
	http.ListenAndServe(":"+args.Port, h)
}

func ServeJSON(e interface{}, w http.ResponseWriter) {
	if e == nil {
		http.Error(w, "No data", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ServeJSONBytes(data, w)
}

func ServeJSONBytes(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	writeAccessOptions(w)
	w.Write(data)
}

func writeAccessOptions(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers",
		"Content-Type, Access-Control-Allow-Headers")
}