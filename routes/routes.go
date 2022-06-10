package routes

import (
	"listcontacts/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/contacts", controllers.GetContacts).Methods("GET")
	router.HandleFunc("/contacts", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/contacts/{id}", controllers.GetContact).Methods("GET")
	router.HandleFunc("/contacts/{id}", controllers.UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id}", controllers.DeleteContact).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
