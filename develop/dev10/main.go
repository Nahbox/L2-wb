package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Обработка аргументов командной строки
	host := flag.String("host", "", "Хост (IP или доменное имя)")
	port := flag.Int("port", 0, "Порт")
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут подключения")
	flag.Parse()

	// Проверка наличия хоста и порта
	if *host == "" || *port == 0 {
		fmt.Println("Необходимо указать хост и порт")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Формирование адреса сервера
	serverAddr := fmt.Sprintf("%s:%d", *host, *port)

	// Подключение к серверу с установленным таймаутом
	conn, err := net.DialTimeout("tcp", serverAddr, *timeout)
	if err != nil {
		fmt.Printf("Ошибка подключения к серверу: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Канал для ожидания сигнала завершения программы
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	// Горутина для чтения данных из сокета и вывода их в STDOUT
	go func() {
		io.Copy(os.Stdout, conn)
		done <- os.Interrupt // Отправка сигнала завершения при закрытии сокета
	}()

	// Горутина для чтения данных из STDIN и записи их в сокет
	go func() {
		io.Copy(conn, os.Stdin)
		done <- os.Interrupt // Отправка сигнала завершения при закрытии STDIN
	}()

	// Ожидание сигнала завершения программы
	<-done
}
