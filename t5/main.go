package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*
утилита grep
*/

//Пример запроса echo "who" | go run main.go -A=2

//name `filePath` (string) | str `search` (string)  | -A `after` - N(int) | -B `before` - N(int) | -C `context` -N(int) | c `count` (bool) | i `ignore-case` (bool) | v `invert` (bool) | F `fixed` (bool) | n `line num`
func main() {
	var name, search, input string
	var A, B, C int
	var c, i, v, F, n bool
	for {
		data := make([]byte, 1)
		n, err := os.Stdin.Read(data)
		if err == nil && n > 0 {
			input += string(data)
		} else {
			break
		}
	}
	search = strings.TrimSpace(input)

	flag.BoolVar(&c, "c", false, "count")
	flag.BoolVar(&i, "i", false, "ignore-case")
	flag.BoolVar(&v, "v", false, "invert") //
	flag.BoolVar(&F, "F", false, "fixed")
	flag.BoolVar(&n, "n", false, "line num") //Номерация строк начинается с 0

	flag.StringVar(&name, "name", "test.txt", "file-path")

	flag.IntVar(&A, "A", 0, "after")
	flag.IntVar(&B, "B", 0, "before")
	flag.IntVar(&C, "C", 0, "context")

	flag.Parse()

	data := readFile(name, i, F)

	if c == true || n == true {
		if c == true {
			fmt.Println(len(data))
		}
		if n == true {
			fmt.Println(lineNum(data, search))
		}
	} else {
		if C != 0 {
			sampleSearchC(data, search, C, v)
		} else if A != 0 && B != 0 {
			sampleSearchC(data, search, A+B, v)
		} else if A != 0 {
			sampleSearchA(data, search, A, v)
		} else if B != 0 {
			sampleSearchB(data, search, B, v)
		} else {
			fmt.Println(data)
		}
	}
}

func readFile(name string, ig, F bool) (data []string) {
	file, err := os.OpenFile(name, os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v\n", err)
		return
	}
	defer file.Close()
	data = make([]string, 0, 4)
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if ig == true {
			if F == true {
				data = append(data, strings.ToLower(line))
			} else {
				data = append(data, strings.TrimSpace(strings.ToLower(line)))
			}
		} else {
			if F == true {
				data = append(data, line)
			} else {
				data = append(data, strings.TrimSpace(line))
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("Ошибка чтения файла: %v\n", err)
				return
			}
		}
	}
	return
}

func lineNum(data []string, search string) (output []int) {
	output = make([]int, 0, 10)
	for i, str := range data {
		if search == str {
			output = append(output, i) //Номерация строк начинается с 0
		}
	}
	fmt.Println(output)
	return
}

func sampleSearchA(data []string, search string, N int, v bool) {
	output := make([]string, 0, 10)

	for i, str := range data {
		if search == str {
			if N < len(data)-(i+1) {
				if v != true {
					for _, val := range data[i+1 : i+N+1] {
						output = append(output, val)
					}
				} else {
					for j := 0; j < N; j++ {
						copy(data[i+1:], data[i+2:])
						data = data[:len(data)-1]
					}
				}
			} else if len(data)-N != 0 {
				if v != true {
					for _, val := range data[i+1:] {
						output = append(output, val)
					}
				} else {
					data = data[:i+1]
				}
			}
		}
	}
	if v != true {
		fmt.Println(output)
	} else {
		fmt.Println(data)
	}
}

func sampleSearchB(data []string, search string, N int, v bool) {
	output := make([]string, 0, 10)
	for i, str := range data {
		if search == str {
			if N < (i + 1) {
				if v != true {
					for _, val := range data[i-N : i] {
						output = append(output, val)
					}
				} else {
					for j := 1; j <= N; j++ {
						copy(data[i-(j+1):], data[i-j:])
						data = data[:len(data)-1]
					}
				}
			} else if N-i != 0 {
				if v != true {
					for _, val := range data[:i] {
						output = append(output, val)
					}
				} else {
					data = data[i:]
				}
			}
		}
	}
	if v == true {
		fmt.Println(data)
	} else {
		fmt.Println(output)
	}
}

func sampleSearchC(data []string, search string, N int, v bool) {

	if v != true {
		output := make([]string, 0, 10)
		for i, str := range data {
			if search == str {
				if N < (i+1) && N < len(data)-(i+1) {
					for _, val := range data[i-N : i] {
						output = append(output, val)
					}
					for _, val := range data[i+1 : i+N+1] {
						output = append(output, val)
					}
				} else if N-i != 0 && N < len(data)-(i+1) {
					for _, val := range data[:i] {
						output = append(output, val)
					}
					for _, val := range data[i+1 : i+N+1] {
						output = append(output, val)
					}
				} else if N < (i+1) && len(data)-N != 0 {
					for _, val := range data[i-N : i] {
						output = append(output, val)
					}
					for _, val := range data[i+1:] {
						output = append(output, val)
					}
				}
			}
		}
		fmt.Println(output)

	} else {
		for i, str := range data {
			if search == str { //ошибка не видит третий who
				if N < (i+1) && N < len(data)-(i+1) {
					var j int
					for j = 1; j <= N; j++ {
						copy(data[i-(j+1):], data[i-j:])
						data = data[:len(data)-1]
					}
					for k := 0; k < N; k++ {
						copy(data[(i+1-j):], data[(i+2-j):])
						data = data[:len(data)-1]
					}
				} else if N-i != 0 && N < len(data)-(i+1) {
					var k int = N - i
					data = data[i:]
					for j := 0; j < N; j++ {
						copy(data[(i+1-k):], data[(i+2-k):])
						data = data[:len(data)-1]
					}

				} else if N < (i+1) && len(data)-N != 0 {
					var j int
					for j = 1; j <= N; j++ {
						copy(data[i-(j+1):], data[i-j:])
						data = data[:len(data)-1]
					}
					data = data[:(i + 1 - j)]
				}
			}
		}
		fmt.Println(data)
	}

}

/*
for i, str := range data {
	if search == str {
		if N < (i + 1) {
			for j := 1; j <= N; j++ {
				copy(data[i-(j+1):], data[i-j:])
				data = data[:len(data)-1]
			}
		} else if N-i != 0 {
			data = data[i:]
		}
	}
}

for i, str := range data {
	if search == str {
		if N < len(data)-(i+1) {
			for j := 0; j < N; j++ {
				copy(data[i+1:], data[i+2:])
				data = data[:len(data)-1]
			}
		} else if len(data)-N != 0 {
			data = data[:i+1]

		}
	}
}
*/
