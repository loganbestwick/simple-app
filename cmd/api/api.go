package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Setup() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/echo", echo)
	router.HandleFunc("/health", health)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func echo(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("value")
	fmt.Fprintf(w, query)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
