package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Healthcheck", "True")
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/healthcheck", healthcheck)
	router.HandleFunc("/weather", getWeather)
	log.Fatal(http.ListenAndServe(":8080", router))
}
