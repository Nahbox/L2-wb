package facadepattern

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// NTPFacade - фасад для работы с NTP
type NTPFacade struct{}

// GetCurrentTime возвращает текущее время
func (f *NTPFacade) GetCurrentTime() (time.Time, error) {
	return ntp.Time("pool.ntp.org")
}

// PrintTime выводит переданное время в консоль
func PrintTime(t time.Time) {
	fmt.Println("Текущее время:", t.Format(time.RFC3339))
}

// Facade : точка входа в паттерн фасад
func Facade() {
	// Создаем экземпляр фасада
	ntpFacade := &NTPFacade{}

	// Получаем текущее время с использованием фасада
	currentTime, err := ntpFacade.GetCurrentTime()
	if err != nil {
		// Обрабатываем ошибку, выводим в STDERR и возвращаем ненулевой код выхода
		fmt.Fprintf(os.Stderr, "Ошибка при получении времени: %v\n", err)
		os.Exit(1)
	}

	// Выводим текущее время
	PrintTime(currentTime)
}
