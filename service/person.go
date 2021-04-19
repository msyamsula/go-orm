package service

import (
	"encoding/json"
	"fmt"
	"go-orm/connection"
	"go-orm/jsonRequest"
	"go-orm/jsonResponse"
	"go-orm/model"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Homelink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("homepage is called")
	fmt.Fprintf(w, "Welcome Home\n")
}

func GetPersonById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get person by id is called")
	personId := mux.Vars(r)["id"]
	db := connection.Db

	var person model.Person
	result := db.First(&person, personId)
	err := result.Error

	// error handling
	if err != nil {
		fmt.Println("error getting a person")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	var res jsonResponse.Person
	res.Name = person.Name
	res.Id = person.ID
	res.Address = person.Address
	res.AccountId = person.AccountId
	json.NewEncoder(w).Encode(res)
}

func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all person is called")
	db := connection.Db

	var persons []model.Person
	result := db.Find(&persons)
	err := result.Error
	if err != nil {
		fmt.Println("error getting people")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var res []jsonResponse.Person
	for _, p := range persons {
		var sres jsonResponse.Person
		sres.AccountId = p.AccountId
		sres.Name = p.Name
		sres.Id = p.ID
		sres.Address = p.Address
		res = append(res, sres)
	}
	json.NewEncoder(w).Encode(res)

}

func AddPerson(w http.ResponseWriter, r *http.Request) {
	db := connection.Db

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error when parsing body")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var body jsonRequest.Person
	err = json.Unmarshal(reqBody, &body)
	if err != nil {
		fmt.Println("error when unmarshall request body")
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	var person model.Person
	person.AccountId = body.AccountId
	person.Name = body.Name

	result := db.Create(&person)
	err = result.Error
	if err != nil {
		fmt.Println("error in database creation")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res jsonResponse.Person
	res.Id = person.ID
	res.AccountId = person.AccountId
	res.Name = person.Name
	res.Address = person.Address

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
