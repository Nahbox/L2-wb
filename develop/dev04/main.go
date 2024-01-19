package main

import (
	"sort"
	"strings"
)

// findAnagrams находит все множества анаграмм в заданном массиве слов
func findAnagrams(words []string) map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range words {
		// Приведем слово к нижнему регистру
		word = strings.ToLower(word)

		// Проверка наличия ключа и добавление в множество анаграмм
		anagramSets = checkAndAddToSet(word, anagramSets)
	}

	// Удаляем множества из одного элемента
	for key, set := range anagramSets {
		if len(set) < 2 {
			delete(anagramSets, key)
		} else {
			// Сортируем множество по возрастанию
			sort.Strings(set)
		}
	}

	return anagramSets
}

// getSortedKey возвращает отсортированный ключ для заданного слова
func checkAndAddToSet(word string, anagramSets map[string][]string) map[string][]string {
	// Преобразуем слово в массив рун для корректной сортировки utf8
	wordRunes := []rune(word)
	keyIsFind := false

	for key, _ := range anagramSets {
		// Сортируем массив рун проверяемого слова
		sort.Slice(wordRunes, func(i, j int) bool {
			return wordRunes[i] < wordRunes[j]
		})

		keyRunes := []rune(key)
		// Сортируем массив рун проверяемого слова
		sort.Slice(keyRunes, func(i, j int) bool {
			return keyRunes[i] < keyRunes[j]
		})

		if string(wordRunes) == string(keyRunes) {
			anagramSets[key] = append(anagramSets[key], word)
			keyIsFind = true
			break
		}
	}

	if !keyIsFind {
		anagramSets[word] = append(anagramSets[word], word)
	}

	// Преобразуем отсортированный массив рун обратно в строку
	return anagramSets
}

func main() {
	// Пример использования функции
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagramSets := findAnagrams(words)

	// Вывод результата
	for key, set := range anagramSets {
		println("Множество анаграмм для ключа", key, ":", strings.Join(set, ", "))
	}
}
