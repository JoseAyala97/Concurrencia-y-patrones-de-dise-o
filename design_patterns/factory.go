package main

import "fmt"

// Patrón de diseño FACTORY
// Clase principal
type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

// constructor para el objeto
func newLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

type Desktop struct {
	Computer
}

// constructor para el objeto
func NewDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 35,
		},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {
	//Validación para identificar el tipo de computador que se va a obtener
	if computerType == "laptop" {
		return newLaptop(), nil
	}
	if computerType == "desktop" {
		return NewDesktop(), nil
	}
	return nil, fmt.Errorf("invalid computer Type")
}

func printNameAndStock(p IProduct) {
	fmt.Printf("Product name: %s, with stock %d\n", p.getName(), p.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")
	printNameAndStock(laptop)
	printNameAndStock(desktop)
}
