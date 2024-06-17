package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	postcardzFile = filepath.Join(filepath.Dir(os.Args[0]), "postcards.json")

	err := readPostcards()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", serveTemplateAdmin).Methods("GET")
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./resources"))))
	// r.HandleFunc("/api/postcards", getPostcards).Methods("GET")
	// r.HandleFunc("/api/postcard/{postcarduuid}", serveTemplateCardForUser).Methods("GET")
	r.HandleFunc("/api/postcard/{postcarduuid}/code", codeForExistingPostcard).Methods("GET")
	// r.HandleFunc("/api/postcard/{postcarduuid}", updatePostcard).Methods("PUT")
	// r.HandleFunc("/api/postcard", createPostcard).Methods("POST")

	c := cors.New(cors.Options{
		AllowedMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost},
	})

	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(":8081", handler))
}

func codeForExistingPostcard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("codeForExistingPostcard with uuid:", vars["postcarduuid"])
	_ = QrCodeOverlay().Render(r.Context(), w)
}

func serveTemplateCardForUser(w http.ResponseWriter, r *http.Request) {
	_ = SiteLayout(false).Render(r.Context(), w)
}

func serveTemplateAdmin(w http.ResponseWriter, r *http.Request) {
	_ = SiteLayout(true).Render(r.Context(), w)
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
