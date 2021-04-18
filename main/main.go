package main

import (
	"go-orm/connection"
	"go-orm/service"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
)

func main() {
	// mysql101.Crud()
	connection.Connect("root", "mysql", "0.0.0.0", "3306", "gorm")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", service.Homelink)
	router.HandleFunc("/person/{id}", service.GetPersonById).Methods("GET")
	router.HandleFunc("/person", service.GetAllPerson).Methods("GET")
	router.HandleFunc("/person", service.GetPersonById).Methods("POST")

	go func() {
		log.Println("Server starting ....")
		err := http.ListenAndServe(":5000", router)
		if err != nil {
			ch <- os.Interrupt
		}
	}()

	<-ch

	log.Println("server stopped")

}
