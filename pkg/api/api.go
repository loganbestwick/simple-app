package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Setup() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func echo(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("value")
	fmt.Fprintf(w, query)
}
