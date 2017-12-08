package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	// appengine.Main()
	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe(":80", nil))
}

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tpl.ExecuteTemplate(w, "index.html", "Hello World!"); err != nil {
		log.Fatal(err)
	}
	// t, _ := template.ParseFiles("index.html")
	// t.Execute(w, "Hello World!")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
