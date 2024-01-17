package facadepattern

import "fmt"

// Подсистема оформления заказа
type OrderSubsystem struct{}

func (os *OrderSubsystem) createOrder() {
	fmt.Println("Заказ успешно создан")
}

// Подсистема проверки инвентаря
type InventorySubsystem struct{}

func (is *InventorySubsystem) checkInventory() {
	fmt.Println("Инвентарь проверен")
}

// Подсистема обработки платежа
type PaymentSubsystem struct{}

func (ps *PaymentSubsystem) processPayment() {
	fmt.Println("Платеж успешно обработан")
}

// Фасад для упрощенного взаимодействия с подсистемой ордера
type OrderFacade struct {
	orderSubsystem     *OrderSubsystem
	inventorySubsystem *InventorySubsystem
	paymentSubsystem   *PaymentSubsystem
}

func NewOrderFacade() *OrderFacade {
	return &OrderFacade{
		orderSubsystem:     &OrderSubsystem{},
		inventorySubsystem: &InventorySubsystem{},
		paymentSubsystem:   &PaymentSubsystem{},
	}
}

// Метод для оформления заказа с использованием фасада
func (of *OrderFacade) PlaceOrder() {
	of.orderSubsystem.createOrder()
	of.inventorySubsystem.checkInventory()
	of.paymentSubsystem.processPayment()
	fmt.Println("Заказ успешно оформлен")
}

// FacadePattern точа входа
func FacadePattern() {
	// Использование фасада для оформления заказа
	orderFacade := NewOrderFacade()
	orderFacade.PlaceOrder()
}
