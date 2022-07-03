package main

import "testing"

func TestUnpack(t *testing.T) {
	//Arrange тестовая выборка
	testTable := []struct {
		input    pkgString
		expected string
	}{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abcd",
			expected: "abcd",
		},
		{
			input:    "45",
			expected: "(некорректная строка)",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "qwerty0y",
			expected: "qwerty",
		},
		{
			input:    `qwe\4\5`,
			expected: "qwe45",
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
	}
	//Act расчет результата
	for _, testCase := range testTable {
		result := pkgString.Unpack(testCase.input)

		//Assert сравнение результата с ожиданием
		if result != testCase.expected {
			t.Errorf("Incorrext result. Expect %s, got %s", testCase.expected, result)
		}

	}
}
