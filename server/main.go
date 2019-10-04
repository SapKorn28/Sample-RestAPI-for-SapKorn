package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// GoDoc

// User user data struct
type User struct {
	ID       string `json:"id"`
	Password string `json:"pass"`
}

func getFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Get Method from Go!"))
}

func postFunc(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("post: ", user)
	json.NewEncoder(w).Encode(user)
}

func putFunc(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("put: ", user)
	json.NewEncoder(w).Encode(user)
}

func deleteFunc(w http.ResponseWriter, r *http.Request) {
	// proceesing...
}

func main() {
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type"})

	router := mux.NewRouter()

	router.HandleFunc("/api", getFunc).Methods("GET")
	router.HandleFunc("/api", postFunc).Methods("POST")
	router.HandleFunc("/api", putFunc).Methods("PUT")
	router.HandleFunc("/api", deleteFunc).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":7070", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}
