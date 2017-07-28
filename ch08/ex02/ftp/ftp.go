package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	defer c.Close()
	for input.Scan() {
		command := strings.Split(input.Text(), " ")
		switch command[0] {
		case "cd":
		case "ls":
		default:
		}
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("ftp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
