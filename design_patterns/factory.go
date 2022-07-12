package main

import "fmt"

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

func (computer *Computer) setStock(stock int) {
	computer.stock = stock
}

func (computer *Computer) setName(name string) {
	computer.name = name
}

func (computer *Computer) getName() string {
	return computer.getName()
}

func (computer *Computer) getStock() int {
	return computer.getStock()
}

type laptop struct {
	Computer
}

func newLaptop() IProduct {
	return &laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

type desktop struct {
	Computer
}

func newDesktop() IProduct {
	return &desktop{
		Computer: Computer{
			name:  "Desktop computer",
			stock: 10,
		},
	}
}

func GetComputerfactory(computerType string) (IProduct, error) {
	if computerType == "laptop" {
		return newLaptop(), nil
	}

	if computerType == "desktop" {
		return newDesktop(), nil
	}

	return nil, fmt.Errorf("Invalid computer type")
}

func printNameAndStock(product IProduct) {
	fmt.Printf("Product name: %s Product stock: %d\n", product.getName(), product.getStock())
}

func main() {
	laptop, _ := GetComputerfactory("laptop")
	desktop, _ := GetComputerfactory("desktop")

	printNameAndStock(laptop)
	printNameAndStock(desktop)
}
