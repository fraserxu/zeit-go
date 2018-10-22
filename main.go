package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
