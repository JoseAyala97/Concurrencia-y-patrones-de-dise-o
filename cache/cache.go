package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFibonacci(n int) int {
	fmt.Printf("calculate Expensive Fibonacc for %d", n)
	time.Sleep(5 * time.Second)
	return n
}

type Service struct {
	//mapa de tipo entero y con condición bool
	InProgress map[int]bool
	//mapa slices de canales
	IsPending map[int][]chan int
	Lock      sync.RWMutex
}

func (s *Service) Work(job int) {
	//Bloquea el programa para poder realizar la lógica de verificar si se esta realizando un trabajo o no
	s.Lock.RLock()
	//Verifica si hay algún trabajo en curso
	exists := s.InProgress[job]
	if exists {
		//Si existe un trabajo en curso, desbloquea nuevamente
		s.Lock.RUnlock()
		//Incluye un bloqueo hasta que este la respuesta
		response := make(chan int)
		defer close(response)
		//Bloquea de nuevo el programa, para agregar a la cola de trabajos pendientes el que quiere iniciarse
		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()
		//Hasta que se realice la finalización de ese trabajo
		fmt.Printf("Waiting for Response job: %d\n", job)
		resp := <-response
		fmt.Printf("Response Done, received %d\n", resp)
		return
	}
	//En caso de que no haya ningún trabajo en curso, desbloquea el programa
	s.Lock.RUnlock()
	//Bloquea nuevamente para poder empezar con la lógica de que no hay trabajo en curso
	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()
	//Empieza a realizar logica para hacer calculo del trabajo
	fmt.Printf("Calculate Fibonacci for %d\n", job)
	result := ExpensiveFibonacci(job)

	s.Lock.RLock()
	pendingWorkers, exists := s.IsPending[job]
	s.Lock.RUnlock()

	if exists {
		for _, pendingWorker := range pendingWorkers {
			//Notificar a todos los workers pendientes que ya el resultado fue enviado
			pendingWorker <- result
		}
		fmt.Printf("Result sent - all pending workers ready job:%d\n", job)
	}
	s.Lock.Lock()
	s.InProgress[job] = false
	//Slices de canales vacío
	s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()
}

// Constructor
func NewServices() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  map[int][]chan int{},
	}
}

func main() {
	service := NewServices()
	jobs := []int{3, 4, 5, 5, 4, 8, 8, 8}
	var wg sync.WaitGroup
	wg.Add(len(jobs))
	for _, n := range jobs {
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(n)
	}
	//Lo ponemos a que espere
	wg.Wait()
}
