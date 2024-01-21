Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
любой порядок чисел от 1 до 8
0
0
.
.
.
```
1. В функции merge не выполняется закрытие канала c, поэтому главная горутина не получит уведомления о закрытии, и цикл for v := range c будет продолжать читать из канала c, даже после того как каналы a и b закроются.
2. Когда закрытый канал c возвращает значения, тип значения по умолчанию (zero value) для типа int - это 0. Поэтому после закрытия канала будут выводиться нули.