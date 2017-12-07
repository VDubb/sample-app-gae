package hello

import (
	"html/template"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, "Hello World!")
}

func init() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}
