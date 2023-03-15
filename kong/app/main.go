package main

import "net/http"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Template(w, "home.html", nil)
}

func main() {
	http.HandleFunc("/", LoginHandler)
	http.HandleFunc("/home", HomeHandler)
	http.ListenAndServe(":8080", nil)
}
