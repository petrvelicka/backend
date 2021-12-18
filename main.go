package main

import (
	"doucovna/doucovna"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var validExtensions = []string{".css", ".png", ".html"}

func isValidFileExtension(path string) bool {
	for _, extension := range validExtensions {
		if strings.HasSuffix(path, extension) {
			return true
		}
	}

	return false
}

var db doucovna.DbConnector

func getTutors(w http.ResponseWriter, r *http.Request) {
	tutors, err := db.GetTutorsJson()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, tutors)
}

func getOffers(w http.ResponseWriter, r *http.Request) {
	offers, err := db.GetOffersJson()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, offers)
}

func getSubjects(w http.ResponseWriter, r *http.Request) {
	subjects, err := db.GetSubjectsJson()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, subjects)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/login.html")
	}
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprintf(w, "Error 404, page not found ")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if isValidFileExtension(r.URL.Path) {
		http.ServeFile(w, r, "static"+string(r.URL.Path))
		return
	}

	if r.URL.Path != "/" {
		pageNotFound(w, r)
		return
	}

	http.ServeFile(w, r, "static/index.html")
}

func main() {
	db = doucovna.NewDbConnector("doucovna.db")
	defer db.Close()

	http.HandleFunc("/get_tutors", getTutors)
	http.HandleFunc("/get_offers", getOffers)
	http.HandleFunc("/get_subjects", getSubjects)

	http.HandleFunc("/login", login)
	http.HandleFunc("/tutors", tutorsHandler)

	http.HandleFunc("/", homeHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
