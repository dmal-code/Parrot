package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
)

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called get")
	keys := r.URL.Query()
	json.NewEncoder(w).Encode(keys)
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called post")
	keys := r.URL.Query()
	headers := r.Header
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	keysJson,_ := json.Marshal(keys);
	headersJson,_ := json.Marshal(headers);

	//var valueMap map[string]interface{}
	valueMap := make(map[string]interface{})
	valueMap["parameters"] = string(keysJson[:])
	valueMap["body"] = string(b[:])
	valueMap["headers"] = string(headersJson[:])

	json.NewEncoder(w).Encode(keysJson[:])
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/get/", get).Methods("GET")
	router.HandleFunc("/post/", post).Methods("POST")
	log.Fatal(http.ListenAndServe(":18000", router))
}
