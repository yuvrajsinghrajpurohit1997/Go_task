package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", GetUser).Methods("GET")
	myRouter.HandleFunc("/search/{id}", SearchUser).Methods("GET")
	myRouter.HandleFunc("/searchgroup/{id}", SearchUserGroup).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main() {
	fmt.Println("Go Task")
	handlerequest()

}
