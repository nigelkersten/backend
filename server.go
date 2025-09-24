package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// editing test

var users []User

func main() {
	portPtr := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")

	portStr := fmt.Sprintf(":%d", *portPtr)
	log.Fatal(http.ListenAndServe(portStr, router))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}
