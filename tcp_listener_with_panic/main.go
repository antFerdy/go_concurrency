package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("error connection")
	}

	log.Println("Started listening port 9000")

	for {
		client, err := conn.Accept()
		if err != nil {
			log.Println("Error accepting connection %w", err)
		}

		go handle(client)
	}
}

func handle(conn net.Conn) {
	defer func() {
		v := recover()
		if v != nil {
			log.Printf("Recovering from panic")
			conn.Close()
		}
	}()

	panic("Error reading connection")
}
