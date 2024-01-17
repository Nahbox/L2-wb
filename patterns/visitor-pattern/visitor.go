package visitorpattern

import "fmt"

// Интерфейс для элементов, которые могут быть посещены
type Animal interface {
	Accept(Visitor)
}

// Конкретная структура элемента - собака
type Dog struct{}

func (d *Dog) Accept(v Visitor) {
	v.VisitDog(d)
}

// Конкретная структура элемента - кошка
type Cat struct{}

func (c *Cat) Accept(v Visitor) {
	v.VisitCat(c)
}

// Интерфейс посетителя
type Visitor interface {
	VisitDog(*Dog)
	VisitCat(*Cat)
}

// Конкретный посетитель - кормильщик
type Feeder struct{}

func (f *Feeder) VisitDog(d *Dog) {
	fmt.Println("Кормим собаку")
}

func (f *Feeder) VisitCat(c *Cat) {
	fmt.Println("Кормим кошку")
}

// Конкретный посетитель - врач
type Veterinarian struct{}

func (v *Veterinarian) VisitDog(d *Dog) {
	fmt.Println("Проводим осмотр собаки")
}

func (v *Veterinarian) VisitCat(c *Cat) {
	fmt.Println("Проводим осмотр кошки")
}

// VisitorPattern точка входа
func VisitorPattern() {
	// Создаем коллекцию животных
	animals := []Animal{&Dog{}, &Cat{}}

	// Создаем посетителей
	feeder := &Feeder{}
	veterinarian := &Veterinarian{}

	// Проходим по коллекции животных и применяем посетителей
	for _, animal := range animals {
		animal.Accept(feeder)
		animal.Accept(veterinarian)
	}
}
