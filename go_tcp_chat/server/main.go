package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	address := "localhost:50880"
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

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Received:", message)
		_, err := fmt.Fprintf(conn, "Echo: %s\n", message)
		if err != nil {
			fmt.Println("Error sending response:", err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from client:", err)
	}
	fmt.Println("Client disconnected:", conn.RemoteAddr())
}
