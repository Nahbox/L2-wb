package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Объявление флагов
	fields := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")

	// Парсинг флагов
	flag.Parse()

	// Чтение данных из STDIN
	scanner := NewInputScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Проверка, содержит ли строка разделитель
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		// Разбивка строки на колонки
		columns := strings.Split(line, *delimiter)

		// Выбор указанных колонок
		if *fields != "" {
			selectedFields := parseFields(*fields)
			for _, field := range selectedFields {
				if field > 0 && field <= len(columns) {
					fmt.Print(columns[field-1])
					if field < len(columns) {
						fmt.Print(*delimiter)
					}
				}
			}
		} else {
			// Вывод всей строки, если не указаны колонки
			fmt.Print(line)
		}

		fmt.Println()
	}
}

// parseFields преобразует строку с номерами колонок в массив чисел
func parseFields(fieldsStr string) []int {
	fields := strings.Split(fieldsStr, ",")
	var result []int
	for _, field := range fields {
		if num, err := strconv.Atoi(field); err == nil {
			result = append(result, num)
		}
	}
	return result
}

// InputScanner обертка для bufio.Scanner с целью облегчения тестирования
type InputScanner struct {
	scanner *bufio.Scanner
}

// NewInputScanner создает новый InputScanner
func NewInputScanner(reader io.Reader) *InputScanner {
	return &InputScanner{scanner: bufio.NewScanner(reader)}
}

// Scan вызывает метод Scan сканнера bufio
func (s *InputScanner) Scan() bool {
	return s.scanner.Scan()
}

// Text вызывает метод Text сканнера bufio
func (s *InputScanner) Text() string {
	return s.scanner.Text()
}
