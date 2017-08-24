package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// control connection
	listener, err := net.Listen("tcp", "localhost:8000")
	// dataListener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start [%v]\n", listener.Addr())
	fmt.Println("Supported command is following: pwd, cd, ls, get, exit")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		log.Printf("Accept [%v]\n", conn.RemoteAddr())
		go handleConn(conn) // handle connections concurrently
	}
}

func handleConn(c net.Conn) {
	// set timeout 30 minutes
	c.SetReadDeadline(time.Now().Add(30 * time.Minute))
	// set maximum request length 128 byte
	request := make([]byte, 128)
	defer c.Close()
	for {
		reqLen, err := c.Read(request)
		if err != nil {
			log.Fatal(err)
		}

		if reqLen == 0 {
			continue
		}
		s := string(request[:reqLen])
		log.Printf("Receive [%v]: %v\n", c.RemoteAddr(), s)
		cmdArgs := strings.Split(s, " ")
		if len(cmdArgs) > 1 {
			exec(c, cmdArgs[0], cmdArgs[1:])
		} else {
			exec(c, cmdArgs[0])
		}
	}
}

// FIXME: last argument...
func exec(conn net.Conn, cmd string, params ...[]string) {
	switch cmd {
	case "pwd":
		// ignore params
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(currentDir)
	case "cd":
		if len(params) == 0 {
			fmt.Println("Missing parameter: Command 'cd' needs destination directory.")
			return
		}
		err := os.Chdir(params[0][0])
		if err != nil {
			log.Fatal(err)
		}
	case "ls":
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		children, err := ioutil.ReadDir(currentDir)
		for _, child := range children {
			fmt.Println(child.Name())
		}
	case "get":
		if len(params) == 0 {
			fmt.Println("Missing parameter: Command 'get' needs filename to download.")
		}
		// TODO: open another port for data connection
		if err := sendFile(params[0][0], conn); err != nil {
			log.Fatal(err)
		}
	case "put":
		// TODO: wait to receive file
		fallthrough
	default:
		log.Printf("Command '%s' not found.\n", cmd)
	}
}

func sendFile(path string, conn net.Conn) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Print("file sending.")
	buf := make([]byte, 1024)
	for {
		buflen, err := file.Read(buf)
		if buflen == 0 {
			conn.Write(nil)
			fmt.Println()
			break
		}
		if err != nil {
			fmt.Println()
			return err
		}
		fmt.Print(".")
		conn.Write(buf[:buflen])
	}
	fmt.Println("complete")
	return nil
}
