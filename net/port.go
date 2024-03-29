package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "url to scan")

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			//revisar los puertos que estan en el rango  --> se espera realizar conexión a alguno de estos puertos
			//Dial -> conectarse al protocolo mencionado
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				//Se crea un continue para que el programa continúe buscando para realizar una conexión
				return
			}
			conn.Close()
			fmt.Printf("Port %d is open \n", port)
		}(i)
	}
	wg.Wait()
}
