package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	// "os"
)

func main() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		log.Panic("dial", err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// status, err := bufio.NewReader(conn).ReadString('\n')

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		log.Panic("reading standard input:", err)
	}
}
