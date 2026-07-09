package cassandra

import (
	"github.com/gocql/gocql"
	"log"
)

var Session *gocql.Session

func InitDb(addr, ks string) {
	log.Println("Cassandra keyspace addr: ", addr)
	var err error
	cluster := gocql.NewCluster(addr)
	cluster.Keyspace = ks
	if Session, err = cluster.CreateSession(); err != nil {
		log.Println(err)
		panic(err)
	}
	log.Println("Cassandra db initialized succussfully")
}
