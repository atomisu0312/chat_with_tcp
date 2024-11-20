package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: client ADDRESS:PORT")
		return
	}

	address := os.Args[1]

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server is listening on", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from client:", err)
			break
		}
		fmt.Println("Received:", message)
		_, err = fmt.Fprintf(conn, "Echo: %s", message)
		if err != nil {
			fmt.Println("Error sending response:", err)
			return
		}
	}
	fmt.Println("Client disconnected:", conn.RemoteAddr())
}
