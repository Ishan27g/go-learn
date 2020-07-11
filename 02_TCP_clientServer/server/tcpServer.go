package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
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
	fmt.Println("Starting Server on port : " + port)

	initServer(port)

}

func initServer(port string) {
	var clientCounter = 0
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	//Handle multiple clients using go-routines
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		clientCounter++
		go connectionHandler(conn, clientCounter)

		fmt.Println("Current connected clients : ", clientCounter)
	}
}

func connectionHandler(conn net.Conn, clientID int) {
	//Stay connected untill client exits
	fmt.Println("connected with clientID ", clientID)
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		recv := strings.TrimSpace(string(data))
		if recv == "exit" || recv == "Exit" {
			break
		}
		fmt.Println("Message from client[", clientID, "] : ")
		fmt.Println(recv)

		//return client Id as every message response
		rsp := strconv.Itoa(clientID) + "\n"

		conn.Write([]byte(string(rsp)))
	}
	fmt.Println("Closing connection with clientID ", clientID)
	conn.Close()
}
