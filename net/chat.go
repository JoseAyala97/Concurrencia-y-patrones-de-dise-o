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
	//clientes que se estan uniendo a nuestro chat
	incomingClient = make(chan Client)
	//Clientes que se estan yendo de nuestro chat
	leavingClient = make(chan Client)
	//mensajes que van a ser transmitidos
	messages = make(chan string)
)

var (
	host = flag.String("h", "localhost", "port")
	port = flag.Int("p", 3090, "port")
)

// Cliente 1 -> server -> HandleConnection(client1)

func HandleConnection(conn net.Conn) {
	//Antes se debe asegurar que la conexión se cierra al manejar la conexión
	defer conn.Close()

	message := make(chan string)
	go MessageWrite(conn, message)
	//clientName := conn.RemoteAddr().String() example: Cliente1:2560
	clientName := conn.RemoteAddr().String()

	message <- fmt.Sprintf("Welcome to the server, your name %s\n", clientName)
	messages <- fmt.Sprintf("New client is here, name %s\n", clientName)
	incomingClient <- message

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s:%s\n", clientName, inputMessage.Text())
	}
	//de esta forma estaremos diciendo que el cleinte abandona el chat
	leavingClient <- message
	messages <- fmt.Sprintf("%s said goodbye", clientName)
}

// Escribe los mensajes que se van recibiendo
func MessageWrite(conn net.Conn, messages <-chan string) {
	for messsage := range messages {
		fmt.Fprintln(conn, messsage)
	}
}

func Broadcast() {
	//cuantos clientes tenemos conectados
	clients := make(map[Client]bool)
	for {
		// multiplexacion de canales
		select {
		//caso en el que se recibe un mensaje
		case message := <-messages:
			// se itera sobre los clientes
			for client := range clients {
				//notificara que el mensaje llego
				client <- message
			}
		// anadimos el nuevo cleinte
		case newClient := <-incomingClient:
			clients[newClient] = true
		case leavingClient := <-leavingClient:
			delete(clients, leavingClient)
			//cerrar el canal porque cliente ha abandonado el canal
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
		go HandleConnection(conn)
	}
}
