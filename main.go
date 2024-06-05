package main

import (
	"log"
	"net/http"
	"text/template"
)

var layout string = `<!DOCTYPE html>
<html>
  <head>
    <script src="/static/htmx.min.js"></script>
    <link rel="stylesheet" type="text/css" href="/static/style.css">
    <title>Hello go-htmlx</title>
  </head>
  <body>
    <div id="content">
	  {{template "content" .}}
	</div>
  </body>
</html>`

var content string = `{{define "content"}}
  <p>Hello go-htmlx</p>
  <button hx-post="/clicked" hx-swap="none">Click me</button>
{{end}}`

var mainTempl, contentTempl *template.Template

func main() {
	fs := http.FileServer(http.Dir("./resources"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/clicked", clicked)

	var err error
	mainTempl, err = template.New("main").Parse(layout)
	if err != nil {
		log.Fatal(err)
	}
	contentTempl, err = template.New("content").Parse(content)
	if err != nil {
		log.Fatal(err)
	}
	_, err = mainTempl.AddParseTree(contentTempl.Name(), contentTempl.Tree)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	mainTempl.ExecuteTemplate(w, "main", nil)
}

func clicked(w http.ResponseWriter, r *http.Request) {
	log.Println("clicked")
}
