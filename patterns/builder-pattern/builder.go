package main

import "fmt"

// Продукт - дом
type House struct {
	Foundation string
	Walls      string
	Roof       string
}

// Интерфейс Строителя
type Builder interface {
	BuildFoundation()
	BuildWalls()
	BuildRoof()
	GetHouse() House
}

// Конкретный строитель для обычного дома
type BasicHouseBuilder struct {
	house House
}

func (b *BasicHouseBuilder) BuildFoundation() {
	b.house.Foundation = "Basic Foundation"
}

func (b *BasicHouseBuilder) BuildWalls() {
	b.house.Walls = "Basic Walls"
}

func (b *BasicHouseBuilder) BuildRoof() {
	b.house.Roof = "Basic Roof"
}

func (b *BasicHouseBuilder) GetHouse() House {
	return b.house
}

// Конкретный строитель для дома с улучшенной крышей
type FancyRoofHouseBuilder struct {
	house House
}

func (f *FancyRoofHouseBuilder) BuildFoundation() {
	f.house.Foundation = "Basic Foundation"
}

func (f *FancyRoofHouseBuilder) BuildWalls() {
	f.house.Walls = "Basic Walls"
}

func (f *FancyRoofHouseBuilder) BuildRoof() {
	f.house.Roof = "Fancy Roof"
}

func (f *FancyRoofHouseBuilder) GetHouse() House {
	return f.house
}

// Директор - управляет процессом строительства
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) ConstructHouse() House {
	d.builder.BuildFoundation()
	d.builder.BuildWalls()
	d.builder.BuildRoof()
	return d.builder.GetHouse()
}

// BuildPattern точка входа
func BuildPattern() {
	// Строитель для обычного дома
	basicHouseBuilder := &BasicHouseBuilder{}
	director := NewDirector(basicHouseBuilder)
	basicHouse := director.ConstructHouse()
	fmt.Println("Basic House:")
	fmt.Println(basicHouse)

	// Строитель для дома с улучшенной крышей
	fancyRoofHouseBuilder := &FancyRoofHouseBuilder{}
	director = NewDirector(fancyRoofHouseBuilder)
	fancyRoofHouse := director.ConstructHouse()
	fmt.Println("\nFancy Roof House:")
	fmt.Println(fancyRoofHouse)
}
