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
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		io.WriteString(conn, "\nwrite something here\n")
		fmt.Fprintln(conn, "some more things here")
		fmt.Fprintf(conn, "%v", "a lot of things")

		li.Close()
	}
}
