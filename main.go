package main

import (
	"bufio"
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
			fmt.Println("Error reading:", err.Error())
			break
		}

		handleCommand(connection, command)
	}
}

func getCommand(connection net.Conn) (string, error) {
	message, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		return "", err
	}

	return string(message), nil
}

func handleCommand(connection net.Conn, command string) {
	res, err := connection.Write([]uint8("+OK\r\n"))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
		return
	}

	fmt.Println("Received command: ", command)
	fmt.Println("Message sent to the client: ", res)
}
