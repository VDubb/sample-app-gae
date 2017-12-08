package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
)

const homepage = '<!doctype html>
<html>
<head>
	<style>
		body {
			background-color: blue;
		}
	</style>
</head>
<body>
	<center><h1><font color=white>
		{{ . }}
	</font></h1></center>
</body>
</html>'

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.Parse(homepage)
	t.Execute(w, "Hello GCP!")
	//fmt.Fprint(w, homepage)
}

func main() {
	http.HandleFunc("/", homepageHandler)
	appengine.Main()
}