package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	result := make(chan interface{})

	go func() {
		defer close(result)

		// Создаем канал для отслеживания состояния каждого done-канала
		done := make(chan struct{})

		// Запускаем горутину для каждого done-канала
		for _, ch := range channels {
			go func(ch <-chan interface{}) {
				select {
				case <-ch:
					// Если канал закрыт, отправляем сигнал в done
					close(done)
				case <-done:
					// Если уже есть сигнал в done, игнорируем
				}
			}(ch)
		}

		// Ожидаем закрытия любого из done-каналов
		<-done
	}()

	return result
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()

	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
