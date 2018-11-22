package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called get")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params)
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called post")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/get/{id}", get).Methods("GET")
	router.HandleFunc("/post/{id}", post).Methods("POST")
	log.Fatal(http.ListenAndServe(":18000", router))
}
