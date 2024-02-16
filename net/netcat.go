package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Que los parametros sean parametrisables
var (
	port = flag.Int("p", 3090, "port")
	host = flag.String("h", "localhost", "host")
)

// host:port
// EScribir -> host:port
// leer ->  host:port

func mian() {
	flag.Parse()
	//Creando - generando conexión
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		//Si hay error se terminará la ejecución del programa
		log.Fatal(err)
	}
	//canales de control - Canales de struct
	done := make(chan struct{})
	//La función (anonima) tendrá que leer todo lo que esta en función y escribirlo en consola
	go func() {
		//escritor consola - lector conexión
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()
	//Conexión actua como si fuera escritor y no lector
	CopyContent(conn, os.Stdin)
	conn.Close()
}

func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}

}
