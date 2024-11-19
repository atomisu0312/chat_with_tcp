package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: client ADDRESS:PORT")
		return
	}
	address := os.Args[1]

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("Error connecting: %v\n", err)
		return
	}
	defer conn.Close()

	// Set TCP_NODELAY
	tcpConn, ok := conn.(*net.TCPConn)
	if ok {
		if err := tcpConn.SetNoDelay(true); err != nil {
			fmt.Printf("Error setting TCP_NODELAY: %v\n", err)
			return
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		sendCommands(conn)
	}()

	go func() {
		defer wg.Done()
		handleReplies(conn)
	}()

	wg.Wait()
}

func sendCommands(conn net.Conn) {
	fmt.Println("Waiting:")
	// 標準入力からのデータを読み取る
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text (Ctrl+D to end):")

	for scanner.Scan() {
		// 読み取ったデータを標準出力に表示
		input := scanner.Text()
		fmt.Printf("Input: %s\n", input)

		// データをTCPストリームに送信
		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Printf("Error sending command: %v\n", err)
			return
		}
	}

	// エラーハンドリング
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from standard input: %v\n", err)
	}
}

func handleReplies(conn net.Conn) {
	// サーバーからの返信を標準出力に表示
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Printf("Received: %s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from server: %v\n", err)
	}
}
