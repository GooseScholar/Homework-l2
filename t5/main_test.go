package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {

	//Arrange
	testTable := []struct {
		input struct {
			name string
			i    bool
			F    bool
		}

		expected []string
	}{
		{
			input: struct {
				name string
				i    bool
				F    bool
			}{name: "test.txt",
				i: false,
				F: false,
			},
			expected: []string{"alex", "Who", "Fill", "kolya", "Stepan", "Kolya", "Ivan", "Fill", "kolya", "Stepan", "Ivan", "Fill", "Who", "kolya", "Stepan", "Ivan", "who", "ivan", "Ivan", "Fill", "kolya", "Stepan", "Fill", "kolya", "Stepan", "Kolya", "Ivan", "who", "alex"},
		},
		{
			input: struct {
				name string
				i    bool
				F    bool
			}{name: "test.txt",
				i: true,
				F: false,
			},
			expected: []string{"alex", "who", "fill", "kolya", "stepan", "kolya", "ivan", "fill", "kolya", "stepan", "ivan", "fill", "who", "kolya", "stepan", "ivan", "who", "ivan", "ivan", "fill", "kolya", "stepan", "fill", "kolya", "stepan", "kolya", "ivan", "who", "alex"},
		},
		{
			input: struct {
				name string
				i    bool
				F    bool
			}{name: "test.txt",
				i: false,
				F: true,
			},
			expected: []string{"alex  ", "Who ", "Fill", "kolya", "Stepan ", "Kolya", "Ivan", "Fill", "kolya", "Stepan", "Ivan", "Fill", "Who ", "kolya", "Stepan", "Ivan ", "who ", "ivan", "Ivan", "Fill", "kolya", "Stepan", "Fill", "kolya", "Stepan ", "Kolya", "Ivan", "who ", "alex "},
		},
		{
			input: struct {
				name string
				i    bool
				F    bool
			}{name: "test.txt",
				i: true,
				F: true,
			},
			expected: []string{"alex  ", "who ", "fill", "kolya", "stepan ", "kolya", "ivan", "fill", "kolya", "stepan", "ivan", "fill", "who ", "kolya", "stepan", "ivan ", "who ", "ivan", "ivan", "fill", "kolya", "stepan", "fill", "kolya", "stepan ", "kolya", "ivan", "who ", "alex "},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := readFile(testCase.input.name, testCase.input.i, testCase.input.F)

		t.Logf("Calling readFile(%v), result %s\n", testCase.input, result)

		//Assert
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %s, got %s", testCase.expected, result))
	}
}

func TestLineNum(t *testing.T) {

	//Arrange
	testTable := []struct {
		input struct {
			data   []string
			search string
		}
		expected []int
	}{
		{
			input: struct {
				data   []string
				search string
			}{
				data:   []string{"alex", "Who", "Fill", "kolya", "Stepan", "Kolya", "Ivan", "Fill", "kolya", "Stepan", "Ivan", "Fill", "Who", "kolya", "Stepan", "Ivan", "who", "ivan", "Ivan", "Fill", "kolya", "Stepan", "Fill", "kolya", "Stepan", "Kolya", "Ivan", "who", "alex"},
				search: "Who",
			},
			expected: []int{1, 12},
		},
		{
			input: struct {
				data   []string
				search string
			}{
				data:   []string{"alex", "Who", "Fill", "kolya", "Stepan", "Kolya", "Ivan", "Fill", "kolya", "Stepan", "Ivan", "Fill", "Who", "kolya", "Stepan", "Ivan", "who", "ivan", "Ivan", "Fill", "kolya", "Stepan", "Fill", "kolya", "Stepan", "Kolya", "Ivan", "who", "alex"},
				search: "Stepan",
			},
			expected: []int{4, 9, 14, 21, 24},
		},
		{
			input: struct {
				data   []string
				search string
			}{
				data:   []string{"alex", "Who", "Fill", "kolya", "Stepan", "Kolya", "Ivan", "Fill", "kolya", "Stepan", "Ivan", "Fill", "Who", "kolya", "Stepan", "Ivan", "who", "ivan", "Ivan", "Fill", "kolya", "Stepan", "Fill", "kolya", "Stepan", "Kolya", "Ivan", "who", "alex"},
				search: "kolya",
			},
			expected: []int{3, 8, 13, 20, 23},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := lineNum(testCase.input.data, testCase.input.search)

		t.Logf("Calling lineNum(%v), result %d\n", testCase.input, result)

		//Assert
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %d, got %d", testCase.expected, result))
	}
}

func TestSampleSearchA(t *testing.T) {
	testTable := []struct {
		input struct {
			data   []string
			search string
			N      int
			v      bool
		}
		expected []string
	}{
		{
			input: struct {
				data   []string
				search string
				N      int
				v      bool
			}{
				data:   []string{"alex", "who", "fill", "kolya", "stepan", "kolya", "ivan", "fill", "kolya", "stepan"},
				search: "alex",
				N:      2,
				v:      false,
			},
			expected: []string{"who", "fill"},
		},
		{
			input: struct {
				data   []string
				search string
				N      int
				v      bool
			}{
				data:   []string{"alex", "who", "fill", "kolya", "stepan", "kolya", "ivan", "fill", "kolya", "stepan"},
				search: "who",
				N:      1,
				v:      true,
			},
			expected: []string{"alex", "who", "kolya", "stepan", "kolya", "ivan", "fill", "kolya", "stepan"},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := sampleSearchA(testCase.input.data, testCase.input.search, testCase.input.N, testCase.input.v)

		t.Logf("Calling sampleSearchA(%v), result %s\n", testCase.input, result)

		//Assert
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %s, got %s", testCase.expected, result))
	}
}

func TestSampleSearchB(t *testing.T) {
	testTable := []struct {
		input struct {
			data   []string
			search string
			N      int
			v      bool
		}
		expected []string
	}{
		{
			input: struct {
				data   []string
				search string
				N      int
				v      bool
			}{
				data:   []string{"alex", "who", "fill", "kolya", "stepan", "kolya", "ivan", "fill", "kolya", "stepan"},
				search: "who",
				N:      2,
				v:      false,
			},
			expected: []string{"alex"},
		},
		{
			input: struct {
				data   []string
				search string
				N      int
				v      bool
			}{
				data:   []string{"alex", "who", "fill", "kolya", "stepan", "kolya", "ivan", "fill", "kolya", "stepan"},
				search: "stepan",
				N:      3,
				v:      true,
			},
			expected: []string{"alex", "stepan", "kolya", "stepan"},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := sampleSearchB(testCase.input.data, testCase.input.search, testCase.input.N, testCase.input.v)

		t.Logf("Calling sampleSearchB(%v), result %s\n", testCase.input, result)

		//Assert
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %s, got %s", testCase.expected, result))
	}
}

func TestSampleSearchC(t *testing.T) {
	testTable := []struct {
		input struct {
			data   []string
			search string
			N      int
			v      bool
		}
		expected []string
	}{
		{
			input: struct {
				data   []string
				search string
				N      int
				v      bool
			}{
				data:   []string{"alex", "who", "fill", "kolya", "stepan", "kolya", "ivan", "fill", "kolya", "stepan"},
				search: "stepan",
				N:      1,
				v:      false,
			},
			expected: []string{"kolya", "kolya", "kolya"},
		},
		{
			input: struct {
				data   []string
				search string
				N      int
				v      bool
			}{
				data:   []string{"alex", "who", "fill", "kolya", "stepan", "kolya", "ivan", "fill", "kolya", "stepan"},
				search: "who",
				N:      2,
				v:      true,
			},
			expected: []string{"who", "stepan", "kolya", "ivan", "fill", "kolya", "stepan"},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := sampleSearchC(testCase.input.data, testCase.input.search, testCase.input.N, testCase.input.v)

		t.Logf("Calling sampleSearchC(%v), result %s\n", testCase.input, result)

		//Assert
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %s, got %s", testCase.expected, result))
	}
}
