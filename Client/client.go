package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)

	if err != nil {
		panic(err)
	}

	fmt.Println("Read", len(bs), "bytes")

	fmt.Println(string(bs))

}
