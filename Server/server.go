package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		fmt.Println("Error is not nil in net.Listen")
		panic(err)

	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error is not nil in ln.Accept")
			panic(err)

		}

		s := fmt.Sprintln("The current time on this server is", time.Now())

		io.WriteString(conn, s)
		fmt.Println("Handled connection from", conn.RemoteAddr())
		//conn.RemoteAddr

		conn.Close()
	}
}
