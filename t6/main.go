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

	//fmt.Println(input)
	//fmt.Println(len(input))
	//fmt.Println("")

	sliceInput := strings.Split(input, "\n")
	//fmt.Println(sliceInput)
	//fmt.Println(len(sliceInput))
	//fmt.Println("")
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
	if f != "" {
		columnsF := make([]int, len(text[0]))
		ff := strings.Split(f, " ")
		for _, val := range ff {
			columnsF = append(columnsF, columns[val])
		}
		fmt.Println(columnsF)
		output := make([][]string, len(columnsF))
		for a, val := range text {
			output[a] = make([]string, len(columnsF))
			for b, valB := range columnsF {
				output[a][b] = val[valB]
			}
		}
		fmt.Println(output)
	} else {
		fmt.Println(text)
	}
}

/*
for a, val := range sliceInput {
	valA := make([]string, len(columnsF))
	for b, valB := range val {
		text[a][b] = valB
	}

}
*/
