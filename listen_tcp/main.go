package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"serve/config"
)

func main() {

	// load service config
	conf := config.LoadConfig()

	// Listen for incoming connections on port 8080
	serve, err := net.Listen("tcp", conf.Host+":"+strconv.Itoa(conf.Port))

	fmt.Println("listening on port: ", conf.Port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("started serving tcp connections")
	// Accept incoming connections and handle them
	//ctx, cancel := context.WithCancel(context.Background())
	for {
		conn, err := serve.Accept()
		fmt.Println("serving tcp connection")
		if err != nil {
			fmt.Println(err)
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
			fmt.Printf("Client disconnected: %v\n", err)
			return
		}

		// why you doing this bro
		trimmedMsg := strings.TrimRight(message, "\n")
		if trimmedMsg == "close connection" ||
			trimmedMsg == "499" {
			handleConnClose(conn)
		} else {
			fmt.Printf("Received message from client: %s\n", message)

			_, err = conn.Write([]byte(strings.ToUpper(message)))

			// can also use //fmt.Fprintf(conn, "echo: %s", strings.ToUpper(message))
			// but for the collect good of understanding.. keep it readable
		}

	}
}

// Close the connection when we're done
func handleConnClose(conn net.Conn) {
	fmt.Println("closing connection for host: ", conn.RemoteAddr())
	conn.Close()
}
