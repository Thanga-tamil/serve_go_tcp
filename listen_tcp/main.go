package main

import (
	"bufio"
	"strings"
	//"context"
	"fmt"
	"net"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {
	// Listen for incoming connections on port 8080
	serve, err := net.Listen("tcp", HOST+":"+PORT)
	fmt.Println("listening on port: ", PORT)
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
	// Close the connection when we're done
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client disconnected: %v\n", err)
			return
		}
		fmt.Printf("Received message from client: %s\n", message)

		_, err = conn.Write([]byte(strings.ToUpper(message)))

		// can also use //fmt.Fprintf(conn, "echo: %s", strings.ToUpper(message))
		// but for the collect good of understanding.. keep it readable

	}
}
