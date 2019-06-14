package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(res http.ResponseWriter, req *http.Request) {
	// io.WriteString(res, "Hello")
	tpl.ExecuteTemplate(res, "index.html", nil)
}

func main() {
	//test
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}
