package main

import "net/http"

func sayHello(w http.ResponseWriter, r *http.Request) {
	planet := r.URL.Query().Get("planet")
	w.Write([]byte("Hello, " + planet))
}

func sayHelloMars(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Mars"))
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/mars", sayHelloMars)
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		panic(err)
	}
}

