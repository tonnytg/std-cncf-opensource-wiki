package main

import "net/http"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func main() {
	http.HandleFunc("/", LoginHandler)
	http.ListenAndServe(":8080", nil)
}
