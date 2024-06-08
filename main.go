package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./resources"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/clicked", clicked)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe("192.168.178.36:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	_ = SiteLayout().Render(r.Context(), w)
}

func clicked(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("clicked with form:", r.PostForm)
}
