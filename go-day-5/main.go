package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(conn, "\nHello from TCP server \n")
		fmt.Fprintln(conn, "How is your day")
		fmt.Fprintln(conn, "%v", "well, I hope")

		conn.Close()
	}
}
