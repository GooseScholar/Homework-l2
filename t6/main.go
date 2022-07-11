package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
Утилита cut
*/

func main() {
	var input, f, d string
	var s bool

	for {
		data := make([]byte, 1)
		n, err := os.Stdin.Read(data)
		if err == nil && n > 0 {
			input += string(data)
		} else {
			break
		}
	}

	flag.BoolVar(&s, "s", false, "separated") //только строки с разделителем

	flag.StringVar(&f, "f", "", "fields")    //выбрать поля (колонки)
	flag.StringVar(&d, "d", "", "delimiter") //использовать другой разделитель

	flag.Parse()

	sliceInput := strings.Split(input, "\n")

	text := make([][]string, len(sliceInput))
	for a, valA := range sliceInput {
		if d != "" {
			val := strings.Split(valA, d)
			text[a] = make([]string, len(val))
			for b, valB := range val {
				text[a][b] = valB
			}
		} else {
			val := strings.Fields(valA)
			text[a] = make([]string, len(val))
			for b, valB := range val {
				text[a][b] = valB
			}

		}

	}

	columns := make(map[string]int)
	for i, val := range text[0] {
		columns[val] = i
	}

	var count int = 0
	if s == true {
		for a, val := range text {
			if len(val) == 1 {
				count++
			} else {
				text[a-count] = text[a]
			}
		}
	}

	if f != "" {
		ff := strings.Split(f, " ")
		columnsF := make([]int, 0, len(ff))
		for _, val := range ff {
			fmt.Println(val)
			columnsF = append(columnsF, columns[val])
		}

		output := make([][]string, len(text)-count)

		for a, _ := range output {
			output[a] = make([]string, 0, len(columnsF))
			if len(text[a]) != len(text[0]) {
				for _, val := range columnsF {
					fmt.Println(val)
					fmt.Println(columnsF)
					if val < len(text[a]) {
						output[a] = append(output[a], text[a][val])
					}
				}
			} else {
				for _, val := range columnsF {
					fmt.Println(text)
					output[a] = append(output[a], text[a][val])
				}
			}
		}

		out := make([]string, 0, len(output))
		for _, val := range output {
			out = append(out, strings.Join(val, " "))
		}
		fmt.Println(strings.TrimSpace(strings.Join(out, "\n")))
	} else {
		out := make([]string, 0, len(text))
		for _, val := range text {
			out = append(out, strings.Join(val, " "))
		}
		fmt.Println(strings.TrimSpace(strings.Join(out, "\n")))
	}

}
