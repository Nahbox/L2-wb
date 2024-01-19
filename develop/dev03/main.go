// Запуск: go run main.go -input input.txt -output output.txt -k 2 -n -r
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Обработка флагов командной строки
	inputFile := flag.String("input", "", "Имя входного файла")
	outputFile := flag.String("output", "", "Имя выходного файла")
	keyColumn := flag.Int("k", 0, "Номер колонки для сортировки (по умолчанию 0 - вся строка)")
	numericSort := flag.Bool("n", false, "Сортировать по числовому значению")
	reverseSort := flag.Bool("r", false, "Сортировать в обратном порядке")
	uniqueSort := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	flag.Parse()

	// Чтение содержимого файла
	content, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		os.Exit(1)
	}

	// Разделение файла на строки
	lines := strings.Split(string(content), "\n")

	// Функция сравнения для сортировки
	comparator := func(i, j int) bool {
		// Выбор колонки для сравнения
		columnI := lines[i]
		columnJ := lines[j]

		if *keyColumn > 0 {
			columnsI := strings.Fields(columnI)
			columnsJ := strings.Fields(columnJ)

			if *keyColumn <= len(columnsI) && *keyColumn <= len(columnsJ) {
				columnI = columnsI[*keyColumn-1]
				columnJ = columnsJ[*keyColumn-1]
			}
		}

		// Применение сортировки
		if *numericSort {
			valI, errI := strconv.Atoi(columnI)
			valJ, errJ := strconv.Atoi(columnJ)
			if errI == nil && errJ == nil {
				return valI < valJ
			}
		}

		return columnI < columnJ
	}

	// Сортировка
	if *uniqueSort {
		sort.SliceStable(lines, func(i, j int) bool {
			return comparator(i, j) && lines[i] != lines[j]
		})
	} else {
		sort.SliceStable(lines, comparator)
	}

	// Если указан флаг для обратной сортировки, переворачиваем строки
	if *reverseSort {
		reverseLines(lines)
	}

	// Запись отсортированных строк в файл
	err = os.WriteFile(*outputFile, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		os.Exit(1)
	}

	fmt.Println("Сортировка успешно завершена.")
}

// Функция для обратной сортировки строк
func reverseLines(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}
