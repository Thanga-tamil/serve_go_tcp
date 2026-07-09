package main

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"

	"serve/cassandra"
	"serve/config"
)

func main() {

	// load service config
	conf := config.LoadConfig()

	dbAddr := conf.Schema.Host + ":" + strconv.Itoa(conf.Schema.Port)
	cassandra.InitDb(dbAddr, conf.Schema.DB)

	tcpAddr := conf.Host + ":" + strconv.Itoa(conf.Port)
	serve, err := net.Listen("tcp", tcpAddr)

	log.Println("serving tcp protocol medium on port: ", conf.Port)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		conn, err := serve.Accept()
		log.Println("Accepted tcp connection for RemoteAddr:", conn.RemoteAddr())
		if err != nil {
			log.Println(err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Client disconnected: %v\n", err)
			return
		}

		// why you doing this bruh
		trimmedMsg := strings.TrimRight(message, "\n")
		if trimmedMsg == "close connection" ||
			trimmedMsg == "499" {
			handleConnClose(conn)
		} else {
			log.Printf("Received message from client: %s\n", message)

			_, err = conn.Write([]byte(strings.ToUpper(message)))

		}

	}
}

// Close the connection when we're done
func handleConnClose(conn net.Conn) {
	log.Println("closing connection for host: ", conn.RemoteAddr())
	conn.Close()
}
