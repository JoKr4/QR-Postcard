package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"net/url"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
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

	r.HandleFunc("/api/postcard/update", updatePostcard).Methods("POST")
	r.HandleFunc("/api/postcard/new", newPostcard).Methods("GET")
	r.HandleFunc("/api/postcard/{postcarduuid}", serveTemplateCardForUser).Methods("GET")
	r.HandleFunc("/api/postcard/{postcarduuid}/code", codeForExistingPostcard).Methods("GET")
	//                                 TODO ?form=code

	c := cors.New(cors.Options{
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
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

type httpQrWriter struct {
	writer http.ResponseWriter
}

func (hqw httpQrWriter) Write(p []byte) (n int, err error) {
	return hqw.writer.Write(p)
}

func (hqw httpQrWriter) Close() error {
	return nil
}

func newPostcard(w http.ResponseWriter, r *http.Request) {
	qrc, err := qrcode.New("https://github.com/yeqown/go-qrcode")
	if err != nil {
		what := fmt.Sprintf("could not generate QRCode: %v", err)
		log.Println(what)
		http.Error(w, what, http.StatusInternalServerError)
		return
	}
	special := httpQrWriter{writer: w}
	options := []standard.ImageOption{
		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
	}
	writer2 := standard.NewWithWriter(special, options...)
	err = qrc.Save(writer2)
	if err != nil {
		what := fmt.Sprintf("could not save QRCode: %v", err)
		log.Println(what)
		http.Error(w, what, http.StatusInternalServerError)
		return
	}
}

func updatePostcard(w http.ResponseWriter, r *http.Request) {

	current, iskey := r.Header["Hx-Current-Url"]
	if !iskey {
		log.Println("in updatePostcard, no such key in request 'Hx-Current-Url'")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if len(current) != 1 {
		log.Println("in updatePostcard, key 'Hx-Current-Url' has more than one element")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	parsed, _ := url.Parse(current[0])
	pathh := parsed.Path
	splitted := strings.Split(pathh, "/")
	if len(splitted) != 4 {
		log.Println("in updatePostcard, key 'Hx-Current-Url' has other than 4 path elements")
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	id := splitted[3]

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	iskey = r.PostForm.Has("usertext")
	if !iskey {
		log.Println("in updatePostcard, key 'usertext' is missing in post form")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	usertext := r.PostForm.Get("usertext")

	ok := false
	for i, p := range postcardz.Postcards {
		if p.UUID == id {
			postcardz.Postcards[i].Textmessage = usertext
			ok = true
			break
		}
	}
	if !ok {
		log.Println("in updatePostcard, the id of postcard to update was not found")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	err = safePostcards()
	if !ok {
		log.Println(err)
		http.Error(w, "could not safe the postcard", http.StatusInternalServerError)
		return
	}

	log.Printf("updated postcard id %s with text %s", id, usertext)

}
