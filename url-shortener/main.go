package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type URL struct {
	ID       int    `json:"id"`
	Original string `json:"original"`
	Short    string `json:"short"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./urls.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS urls (id INTEGER PRIMARY KEY, original TEXT, short TEXT);`)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/shorten", shortenURL).Methods("POST")
	r.HandleFunc("/{short}", redirectURL).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
	var url URL
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url.Short = generateShortURL()

	_, err = db.Exec(`INSERT INTO urls (original, short) VALUES (?, ?)`, url.Original, url.Short)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(url)
}

func redirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	short := vars["short"]

	var url URL
	err := db.QueryRow(`SELECT original FROM urls WHERE short = ?`, short).Scan(&url.Original)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.Original, http.StatusFound)
}

func generateShortURL() string {
	// This is a placeholder function. Implement your own URL shortening logic.
	return "short-url"
}
