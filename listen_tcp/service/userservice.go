package service

import (
	"github.com/gocql/gocql"
	"log"
	"serve/repo"
)

func RegisterUser() {

	id, _ := gocql.RandomUUID()
	id1, _ := gocql.RandomUUID()

	log.Println("google's uuid: ", id)
	repo.InsertUser(id, id1.String(), "damn@gmail.com")

	log.Println("RegisterUser func call completed")

}
