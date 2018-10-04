package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func dateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(time.Now().Local())
}

func main() {
	http.HandleFunc("/", dateHandler)

	http.ListenAndServe(":8080", nil)
}
