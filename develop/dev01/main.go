package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Получаем текущее системное время
	currentTime := time.Now()

	// Получаем точное время с использованием NTP
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при получении времени: %v\n", err)
		os.Exit(1)
	}

	// Выводим текущее и точное время
	fmt.Printf("Текущее время: %s\n", currentTime.Format(time.RFC3339))
	fmt.Printf("Точное время (с использованием NTP): %s\n", ntpTime.Format(time.RFC3339))
}
