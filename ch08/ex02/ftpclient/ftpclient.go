package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connect: %v\n", conn.RemoteAddr())
	defer conn.Close()

	for {
		fmt.Print("ftp > ")
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		switch {
		case stdin.Text() == "exit":
			break
		case strings.HasPrefix(stdin.Text(), "get"):
			conn.Write([]byte(stdin.Text()))
			filename := strings.Split(stdin.Text(), " ")[1]
			err := download(filename, conn)
			if err != nil {
				log.Fatal(err)
			}
		default:
			conn.Write([]byte(stdin.Text()))
		}
	}
}

func download(filename string, conn net.Conn) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Print("file downloading.")
	buf := make([]byte, 1024)
	for {
		buflen, err := conn.Read(buf)
		if buflen == 0 {
			fmt.Println()
			break
		}
		if err != nil {
			fmt.Println()
			return err
		}
		fmt.Print(".")
		file.Write(buf[:buflen])
	}
	fmt.Println("complete")
	return nil
}
