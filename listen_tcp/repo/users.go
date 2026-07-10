package repo

import (
	"log"
	"serve/cassandra"
)

// type SingleEventInsertJob struct {
// 	SensorColour string
// 	TimeBucket   string
// 	EventId      string
// 	Timepoint    time.Time
// }

func InsertUser(user_id, email_id string) {
	q := "INSERT INTO users(user_id, email_id) VALUES (?, ?)"

	query := cassandra.Session.Query(q)

	err := query.Bind(user_id, email_id).Exec()

	if err != nil {
		log.Println("Error while executing query: ", q)
		log.Println("Error: ", err)
	}

}
