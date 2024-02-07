package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 0; i < 100; i++ {
		//revisar los puertos que estan en el rango  --> se espera realizar conexión a alguno de estos puertos
		//Dial -> conectarse al protocolo mencionado
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))
		if err != nil {
			//Se crea un continue para que el programa continúe buscando para realizar una conexión
			continue
		}
		conn.Close()
		fmt.Printf("Port %d is open \n", i)
	}
}
