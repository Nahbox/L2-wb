package builderpattern

import "testing"

func TestBuildString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\45", `qwe44444`},
		{"qwe\\\\5", `qwe\\\\\`},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			director := StringBuildDirector{}
			result := director.BuildString(test.input)

			if result != test.expected {
				t.Errorf("Неправильный результат. Ожидалось %s, получено %s", test.expected, result)
			}
		})
	}
}
