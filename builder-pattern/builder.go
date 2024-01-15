package builderpattern

import (
	"strconv"
)

// StringBuilder - строитель строки
type StringBuilder struct {
	result string
}

// StringBuildDirector - директор для построения строки
type StringBuildDirector struct {
	builder StringBuilder
}

// BuildString - основной метод для построения строки
func (d *StringBuildDirector) BuildString(input string) string {
	d.builder.result = ""
	for i := 0; i < len(input); i++ {
		if input[i] == '\\' {
			i++
			d.handleEscapeSequence(input[i])
		} else if isDigit(input[i]) {
			count, _ := strconv.Atoi(string(input[i]))
			d.handleRepeat(count)
		} else {
			d.handleChar(input[i])
		}
	}
	return d.builder.result
}

// handleEscapeSequence - обработка escape-последовательности
func (d *StringBuildDirector) handleEscapeSequence(ch byte) {
	d.builder.result += string(ch)
}

// handleRepeat - обработка повтора символа
func (d *StringBuildDirector) handleRepeat(count int) {
	if len(d.builder.result) == 0 {
		return
	}
	lastChar := d.builder.result[len(d.builder.result)-1]
	for i := 1; i < count; i++ {
		d.builder.result += string(lastChar)
	}
}

// handleChar - обработка обычного символа
func (d *StringBuildDirector) handleChar(ch byte) {
	d.builder.result += string(ch)
}

// isDigit - проверка, является ли символ цифрой
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// Builder : точка входа в паттерн строитель
func Builder() {
	director := StringBuildDirector{}

	result := director.BuildString("a4bc2d5e")
	println(result)

	result = director.BuildString("abcd")
	println(result)

	result = director.BuildString("45")
	println(result)

	result = director.BuildString("qwe\\4\\5")
	println(result)

	result = director.BuildString("qwe\\45")
	println(result)

	result = director.BuildString("qwe\\\\5")
	println(result)
}
