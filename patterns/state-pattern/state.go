package statepattern

import "fmt"

// Интерфейс состояния
type OrderState interface {
	ProcessOrder() string
}

// Конкретное состояние - заказ оформлен
type OrderPlacedState struct{}

func (ops *OrderPlacedState) ProcessOrder() string {
	return "Order is placed"
}

// Конкретное состояние - заказ отправлен
type OrderShippedState struct{}

func (oss *OrderShippedState) ProcessOrder() string {
	return "Order is shipped"
}

// Контекст - заказ
type OrderContext struct {
	state OrderState
}

func (oc *OrderContext) SetState(state OrderState) {
	oc.state = state
}

func (oc *OrderContext) ProcessOrder() string {
	return oc.state.ProcessOrder()
}

// StatePattern точка входа
func StatePattern() {
	// Создаем заказ и устанавливаем начальное состояние (заказ оформлен)
	order := &OrderContext{state: &OrderPlacedState{}}

	// Обрабатываем заказ в текущем состоянии
	fmt.Println(order.ProcessOrder())

	// Меняем состояние на "заказ отправлен"
	order.SetState(&OrderShippedState{})

	// Обрабатываем заказ в новом состоянии
	fmt.Println(order.ProcessOrder())
}
