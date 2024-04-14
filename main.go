package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":6379")

	if err != nil {
		fmt.Errorf("Fail to listen: %v", err)
		os.Exit(1)
	}

	fmt.Println("Listening on :6379")

	for {
		connection, err := listener.Accept()

		if err != nil {
			fmt.Errorf("Fail to accept connection: %v", err)
			continue
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	fmt.Println("Handle connection")

	// Loop on connection to handle command
	for {
		command, err := getCommand(connection)
		if err != nil {
			connection.Write([]byte{})
      continue
		}

    handleCommand(command)
	}
}

func getCommand(connection net.Conn) (string, error) {
	// TODO: get command from connection
	return "GET abc", nil
}

func handleCommand(command string) {
	// TODO: handle command
}
