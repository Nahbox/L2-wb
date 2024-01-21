Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
deadlock

```
Произойдет блокировка (deadlock) из-за того, что горутина main ожидает чтения из канала ch, но нет закрытия канала после отправки значений в анонимной горутине.