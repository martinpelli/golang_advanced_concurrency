package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingClients         = make(chan Client)
	leavingClients          = make(chan Client)
	messagesToServerChannel = make(chan string)
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.Int("p", 3090, "port number")
)

func Handleconnection(conn net.Conn) {
	defer conn.Close()
	messagesToClientChannel := make(chan string)
	go MessageWriter(conn, messagesToClientChannel)
	clientName := conn.RemoteAddr().String()
	messagesToClientChannel <- fmt.Sprintf("Welcome to the server %s\n", clientName)
	messagesToServerChannel <- fmt.Sprintf("New client entered the server %s\n", clientName)
	incomingClients <- messagesToClientChannel

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messagesToServerChannel <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}

	leavingClients <- messagesToClientChannel
	messagesToServerChannel <- fmt.Sprintf("%s left the server\n", clientName)

}

func MessageWriter(conn net.Conn, messagesToServerChannel <-chan string) {
	for message := range messagesToServerChannel {
		fmt.Fprintf(conn, message)
	}
}

func Broadcast() {
	clients := make(map[Client]bool)
	for {
		select {
		case message := <-messagesToServerChannel:
			for client := range clients {
				client <- message
			}
		case newClient := <-incomingClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}

	go Broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go Handleconnection(conn)
	}
}
