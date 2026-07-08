package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Close the connection
	//defer conn.Close()

	// Send some data to the server
	// _, err = conn.Write([]byte("Hello from client"))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		//if 499
		fmt.Print("Send a message: ")
		message, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(message)) // fmt.Fprintf(conn, message)

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Server error: %v\n", err)
			return
		}
		fmt.Printf("Server says: %s", response)
	}

}
