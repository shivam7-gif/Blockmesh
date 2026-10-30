package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:3000")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Connected to Peer !")

	_, err = conn.Write([]byte("Ping from peer\n"))

	if err != nil {
		log.Fatal(err)
	}
	reply, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Reply , ", reply)
}
