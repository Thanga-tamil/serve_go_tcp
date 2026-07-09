package main

import (
	"bufio"
	"dail/config"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {

	// local config json
	conf := config.LoadConfig()
	// Connect to the server

	conn, err := net.Dial("tcp", conf.Host+":"+strconv.Itoa(conf.Port))
	if err != nil {
		fmt.Println(err)
		return
	}

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
