package main

import "strconv"

// BuildString - основной метод для построения строки
func BuildString(input string) string {
	result := ""
	for i := 0; i < len(input); i++ {
		if input[i] == '\\' {
			i++
			result += string(input[i])
		} else if isDigit(input[i]) {
			count, _ := strconv.Atoi(string(input[i]))
			if len(result) != 0 {
				lastChar := result[len(result)-1]
				for i := 1; i < count; i++ {
					result += string(lastChar)
				}
			}
		} else {
			result += string(input[i])
		}
	}
	return result
}

// isDigit - проверка, является ли символ цифрой
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func main() {
	result := BuildString("a4bc2d5e")
	println(result)

	result = BuildString("abcd")
	println(result)

	result = BuildString("45")
	println(result)

	result = BuildString("qwe\\4\\5")
	println(result)

	result = BuildString("qwe\\45")
	println(result)

	result = BuildString("qwe\\\\5")
	println(result)
}
