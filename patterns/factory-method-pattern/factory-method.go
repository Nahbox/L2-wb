package factorymethodpattern

import (
	"errors"
	"fmt"
)

// Интерфейс для транспорта
type Transport interface {
	Drive() string
}

// Конкретная реализация транспорта - автомобиль
type Car struct{}

func (c *Car) Drive() string {
	return "Driving a car"
}

// Конкретная реализация транспорта - самолет
type Plane struct{}

func (p *Plane) Drive() string {
	return "Flying a plane"
}

// Интерфейс для фабричного метода (TransportFactory)
type TransportFactory interface {
	CreateTransport() Transport
}

// Конкретная реализация фабричного метода для создания автомобиля
type CarFactory struct{}

func (cf *CarFactory) CreateTransport() Transport {
	return &Car{}
}

// Конкретная реализация фабричного метода для создания самолета
type PlaneFactory struct{}

func (pf *PlaneFactory) CreateTransport() Transport {
	return &Plane{}
}

// Фабрика
func GetTransport(transport string) (Transport, error) {
	if transport == "Car" {
		tf := CarFactory{}
		return tf.CreateTransport(), nil
	}
	if transport == "Plane" {
		pf := PlaneFactory{}
		return pf.CreateTransport(), nil
	}
	return nil, errors.New("неправильно указан тип транспорта")
}

// FactoryMethodPattern точка входа
func FactoryMethodPattern() {
	// Используем фабрику для создания автомобиля
	car, _ := GetTransport("Car")
	fmt.Println("Driving a created transport:", car.Drive())

	// Используем фабрику для создания самолета
	plane, _ := GetTransport("Plane")
	fmt.Println("Driving a created transport:", plane.Drive())
}
