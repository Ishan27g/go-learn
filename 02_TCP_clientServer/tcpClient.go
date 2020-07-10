package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter port number")
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Invalid input, Enter valid port number")
		return
	}

	port := ":"
	port += args[1]

	fmt.Println("Connecting client to port : " + port)

	initClient(port)
}

func initClient(port string) {
	conn, err := net.Dial("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	//send dummmy message to get clientID
	fmt.Fprintf(conn, "hello\n")

	//Read response
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	recv := strings.TrimSpace(string(data))

	fmt.Println("clientId : " + recv)

	//Take user input as message
	for {
		fmt.Println("Enter message followed by Enter-key")
		input := bufio.NewReader(os.Stdin)
		msg, _ := input.ReadString('\n')

		//Send message
		fmt.Fprintf(conn, msg+"\n")

		//read response
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		recv := strings.TrimSpace(string(data))
		fmt.Println("Server Response : " + recv)

		//check exit condition
		msg = strings.TrimSpace(string(msg))
		if msg == "exit" || msg == "Exit" {
			fmt.Println("Client exiting...")
			return
		}
	}

	//send exit message to get clientID
	fmt.Fprintf(conn, "exit")
}
