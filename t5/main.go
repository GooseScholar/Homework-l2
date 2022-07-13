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

	flag.BoolVar(&c, "c", false, "count")       // количество строк
	flag.BoolVar(&i, "i", false, "ignore-case") //игнорировать регистр
	flag.BoolVar(&v, "v", false, "invert")      //вместо совпадения, исключать
	flag.BoolVar(&F, "F", false, "fixed")       //точное совпадение со строкой, не паттерн
	flag.BoolVar(&n, "n", false, "line num")    //напечатать номер строки, номерация строк начинается с 0

	flag.StringVar(&name, "name", "test.txt", "file-path")

	flag.IntVar(&A, "A", 0, "after")   //печатать +N строк после совпадения
	flag.IntVar(&B, "B", 0, "before")  //печатать +N строк до совпадения
	flag.IntVar(&C, "C", 0, "context") //печатать ±N строк вокруг совпадения

	flag.Parse()

	data := readFile(name, i, F)

	if c == true || n == true {
		if c == true {
			fmt.Println(len(data))
		}
		if n == true {
			lineNum(data, search)
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

//Чтение файла
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

//Вывод номеров строк совпадения
func lineNum(data []string, search string) {
	out := make([]int, 0, 10)
	for i, str := range data {
		if search == str {
			out = append(out, i) //Номерация строк начинается с 0
		}
	}
	fmt.Println(out)

}

//Вывод N строк после совпадения
func sampleSearchA(data []string, search string, N int, v bool) {
	outB := make([]bool, 0, len(data)) //слайс для учета N строк после совмадения
	for i := 0; i < len(data); i++ {
		outB = append(outB, false)
	}
	out := make([]string, 0, 10)

	for a, str := range data { //поиск совпадений
		if search == str {
			if a+N < len(data)-1 { //если N не выходит за границы слайса
				for b := 1; b <= N; b++ {
					outB[a+b] = true //строки после совпадения заносятся в []bool со значением true
				}
			} else { // N выходит за границы слайса
				for b := 1; b <= len(data)-1-a; b++ {
					outB[a+b] = true //строки после совпадения заносятся в []bool со значением true
				}
			}
		}
	}

	if v != true { //Выводим N строк после совпадения
		for a, str := range outB {
			if str == true {
				out = append(out, data[a])
			}
		}
	} else { //Выводим все строки, кроме N строк после совпадения
		for a, str := range outB {
			if str != true {
				out = append(out, data[a])
			}
		}
	}
	fmt.Println(out)

}

//Вывод N строк перед совпадением
func sampleSearchB(data []string, search string, N int, v bool) {
	outB := make([]bool, 0, len(data)) //слайс для учета N строк перед совмадения
	for i := 0; i < len(data); i++ {
		outB = append(outB, false)
	}
	out := make([]string, 0, 10)

	for a, str := range data { //поиск совпадений
		if search == str {
			if a-N > 0 { //если N не выходит за границы слайса
				for b := 1; b <= N; b++ {
					outB[a-b] = true //строки перед совпадением заносятся в []bool со значением true
				}
			} else { // N вызодит за границы слайса
				for b := 1; b <= a; b++ {
					outB[a-b] = true //строки перед совпадением заносятся в []bool со значением true
				}
			}
		}
	}

	if v != true { //Выводим N строк перед совпадением
		for a, str := range outB {
			if str == true {
				out = append(out, data[a])
			}
		}
	} else { //Выводим все строки, кроме N строк перед совпадением
		for a, str := range outB {
			if str != true {
				out = append(out, data[a])
			}
		}
	}
	fmt.Println(out)
}

//Вывод N строк перед и после совпадением
func sampleSearchC(data []string, search string, N int, v bool) {
	outB := make([]bool, 0, len(data)) //слайс для учета N строк перед/после совмадения
	for i := 0; i < len(data); i++ {
		outB = append(outB, false)
	}
	out := make([]string, 0, 10)

	for a, str := range data { //поиск совпадений
		if search == str { //Разметка N строк перед и после совпадения
			if a-N > 0 { //если N не выходит за границы слайса
				for b := 1; b <= N; b++ {
					outB[a-b] = true //строки перед совпадением заносятся в []bool со значением true
				}
			} else { // N вызодит за границы слайса
				for b := 1; b <= a; b++ {
					outB[a-b] = true //строки перед совпадением заносятся в []bool со значением true
				}
			}

			if a+N < len(data)-1 { //если N не выходит за границы слайса
				for b := 1; b <= N; b++ {
					outB[a+b] = true //строки после совпадения заносятся в []bool со значением true
				}
			} else { // N выходит за границы слайса
				for b := 1; b <= len(data)-1-a; b++ {
					outB[a+b] = true //строки после совпадения заносятся в []bool со значением true
				}
			}
		}

	}
	fmt.Println(outB)
	if v != true { //Выводим N строк перед совпадением
		for a, str := range outB {
			if str == true {
				out = append(out, data[a])
			}
		}
	} else { //Выводим все строки, кроме N строк перед совпадением
		for a, str := range outB {
			if str != true {
				out = append(out, data[a])
			}
		}
	}
	fmt.Println(out)
}
