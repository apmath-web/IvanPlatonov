package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func dateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(time.Now().Local())
}

func waitHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	res := "Nothing to wait"

	strVal := r.URL.Query().Get("delay")
	if _, err := strconv.Atoi(strVal); err == nil {
		dur, _ := time.ParseDuration(strVal + "ms")
		time.Sleep(dur)
		res = "Sleep for " + dur.String() + " completed"
	}

	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/date", dateHandler)
	http.HandleFunc("/wait", waitHandler)

	http.ListenAndServe(":8080", nil)
}
