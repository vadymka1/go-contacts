package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-crud/app"
	"github.com/go-crud/controllers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)

	log.Println("Listening on port ", port)

	if err != nil {
		fmt.Print(err)
	}
}
