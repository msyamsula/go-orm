package mysql101

import (
	"fmt"
	"go-orm/connection"
	"go-orm/model"
	"log"
)

func Crud() {
	connection.Connect("root", "mysql", "0.0.0.0", "3306", "gorm")
	db := connection.Db

	// create
	newPerson := model.Person{
		ID:        10,
		Name:      "Error",
		Address:   "test",
		AccountId: 11234,
	}

	result := db.Create(&newPerson)
	err := result.Error
	if err != nil {
		fmt.Println("error when creating new person")
		fmt.Println(err)
	}

	// get by id
	var person model.Person
	result = db.First(&person, 1)
	if result.Error != nil {
		log.Println("failed to query")
		log.Println(result.Error)
	} else {
		fmt.Println(person)
	}

	// get all
	var persons []model.Person
	result = db.Find(&persons)
	if result.Error != nil {
		log.Println("failed to get all")
		log.Println(result.Error)
	} else {
		for _, p := range persons {
			fmt.Println(p)
		}
	}
}
