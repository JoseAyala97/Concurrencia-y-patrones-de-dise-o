package main

import "fmt"

type Topic interface {
	register(observer Observer)
	//Notificará a los objetos
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

// Item -> no disponible
// Item -> Avise que hay items

type Item struct {
	observers []Observer
	name      string
	//Será quién nos dirá si hay items disponibles o no
	available bool
}

// constructor of item
func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

// Method
func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available \n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) register(observer Observer) {
	//append element to observer
	i.observers = append(i.observers, observer)
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

// create one observer
type EmailClient struct {
	id string
}

func (eC *EmailClient) getId() string {
	return eC.id
}

func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email - %s available from client %s \n", value, eC.id)
}

func main() {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailClient{
		id: "12ab",
	}
	secondObserver := &EmailClient{
		id: "34dc",
	}
	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.UpdateAvailable()
}
