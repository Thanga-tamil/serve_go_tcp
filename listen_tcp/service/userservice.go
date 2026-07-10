package service

import (
	"log"
	"serve/repo"

	uuidv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/request_id/uuid/v3"
)

func RegisterUser() {

	uuid := "f672f0d2-eebe-4cea-8f5a-e9454e840e79"
	uuidv3.new(type)
	log.Println("google's uuid: ", id)
	repo.InsertUser(uuid, "damn@gmail.com")

	log.Println("RegisterUser func call completed")
}
