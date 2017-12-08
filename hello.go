package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, "Hello World!")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe(":80", nil))
}
