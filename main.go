package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/plain")
		w.Write([]byte("Hello World!!!"))
	})
	http.ListenAndServe(":8080", nil)
}
