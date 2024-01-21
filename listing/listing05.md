Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```

1. Функция test() возвращает nil типа *customError.
2. В функции main переменной err присваивается результат вызова test(), который равен nil.
3. Затем программа проверяет, является ли err равным nil. Условие err != nil выполняется, так как err имеет тип error (интерфейс), и сравнение с nil в этом случае идет по значению, а не по указателю. Таким образом, условие считается истинным, и программа выводит "error".