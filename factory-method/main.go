package main

import (
	"fmt"
	"log"
)

/*
Фабричный метод
*/

// action позволяет клиентам узнать доступные действия
type action string

const (
	A action = "A"
	B action = "B"
	C action = "C"
)

// Creator предоставляет фабричный интерфейс
type Creator interface {
	CreateProduct(action action) Product // Фибричный метод
}

// Product предоставляет интерфейс продуктаю. Все продукты возвращаемые фабрикой должны иметь единый интерфейс
type Product interface {
	Use() string
}

// ConcreteCreator реализует интерфейс креатор
type ConcreteCreator struct{}

// ConcreteCreator конструктор.
func NewCreator() Creator {
	return &ConcreteCreator{}
}

// CreateProduct  фабричный метод.
func (p *ConcreteCreator) CreateProduct(action action) Product {
	var product Product

	switch action {
	case A:
		product = &ConcreteProductA{string(action)}
	case B:
		product = &ConcreteProductB{string(action)}
	case C:
		product = &ConcreteProductC{string(action)}
	default:
		log.Fatalln("Unknown Action")
	}

	return product
}

//Реализует продукт "A".
type ConcreteProductA struct {
	action string
}

// Возваращает product action.
func (p *ConcreteProductA) Use() string {
	return p.action
}

//Реализует продукт "B".
type ConcreteProductB struct {
	action string
}

// Возваращает product action.
func (p *ConcreteProductB) Use() string {
	return p.action
}

// Реализует продукт "C".
type ConcreteProductC struct {
	action string
}

// Возваращает product action.
func (p *ConcreteProductC) Use() string {
	return p.action
}

func main() {
	s := []string{"A", "B", "C"}

	factory := NewCreator()

	products := []Product{factory.CreateProduct(A),
		factory.CreateProduct(B),
		factory.CreateProduct(C)}

	fmt.Println(products[0].Use())
	fmt.Println(products[1].Use())
	fmt.Println(products[2].Use())
	fmt.Println(s)
}
