package corpattern

import "fmt"

// Запрос на отпуск
type LeaveRequest struct {
	Name     string
	Duration int
}

// Интерфейс обработчика
type Handler interface {
	HandleRequest(leaveRequest LeaveRequest)
	SetNextHandler(nextHandler Handler)
}

// Конкретный обработчик - Руководитель отдела
type DepartmentManager struct {
	nextHandler Handler
}

func (dm *DepartmentManager) HandleRequest(leaveRequest LeaveRequest) {
	if leaveRequest.Duration <= 5 {
		fmt.Printf("Заявка на отпуск для %s одобрена руководителем отдела.\n", leaveRequest.Name)
	} else if dm.nextHandler != nil {
		fmt.Println("Заявка передана выше.")
		dm.nextHandler.HandleRequest(leaveRequest)
	} else {
		fmt.Println("Нет подходящего уровня руководства для рассмотрения заявки.")
	}
}

func (dm *DepartmentManager) SetNextHandler(nextHandler Handler) {
	dm.nextHandler = nextHandler
}

// Конкретный обработчик - Генеральный директор
type CEO struct {
	nextHandler Handler
}

func (ceo *CEO) HandleRequest(leaveRequest LeaveRequest) {
	if leaveRequest.Duration > 5 {
		fmt.Printf("Заявка на отпуск для %s одобрена генеральным директором.\n", leaveRequest.Name)
	} else if ceo.nextHandler != nil {
		fmt.Println("Заявка передана выше.")
		ceo.nextHandler.HandleRequest(leaveRequest)
	} else {
		fmt.Println("Нет подходящего уровня руководства для рассмотрения заявки.")
	}
}

func (ceo *CEO) SetNextHandler(nextHandler Handler) {
	ceo.nextHandler = nextHandler
}

// CORPattern точка входа
func CORPattern() {
	// Создаем обработчиков
	departmentManager := &DepartmentManager{}
	ceo := &CEO{}

	// Настраиваем цепочку обработчиков
	departmentManager.SetNextHandler(ceo)

	// Создаем запросы на отпуск
	leaveRequest1 := LeaveRequest{Name: "John", Duration: 3}
	leaveRequest2 := LeaveRequest{Name: "Alice", Duration: 7}

	// Обрабатываем запросы с помощью цепочки обработчиков
	departmentManager.HandleRequest(leaveRequest1)
	fmt.Println()
	departmentManager.HandleRequest(leaveRequest2)
}
