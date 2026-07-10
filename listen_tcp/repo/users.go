package repo

import (
	"log"
	"serve/cassandra"
	"time"

	"github.com/gocql/gocql"
)

func InsertUser(id gocql.UUID, user_id, email_id string) {
	q := getQuery()
	query := cassandra.Session.Query(q)

	err := query.Bind(id,
		user_id,
		email_id,
		false,
		"919025565212",
		time.Now(),
	).Exec()

	if err != nil {
		log.Println("Error while executing query: ", q)
		log.Println("Error: ", err)
	}

}

func getQuery() string {
	return "INSERT INTO users(id, user_id, email_id, is_deleted," +
		" mobile_number, created_at) VALUES (?, ?, ?, ?, ?, ?)"
}
