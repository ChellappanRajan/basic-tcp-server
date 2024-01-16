package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, error := net.Listen("tcp", ":8080")

	if error != nil {
		fmt.Println(error)
		return
	}

	defer listener.Close()

	fmt.Println("TCP Server is running on port 8080....")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println((err))
			return
		}
		fmt.Println("Message Received:", string(message))

		conn.Write([]byte("Message received. \n"))
	}

}
