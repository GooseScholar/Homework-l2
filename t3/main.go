package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
утилита sort
*/

//name - имя файла, n - сортировать по числовому значению(0-нет, число - есть, default 0), k - номер колонки (0,1,2,3 и так далее default 0), o - порядок сортировки (0-в порядке возрастания, 1 - в порядке убывания, default 0)
func main() {
	var name string
	var k int
	var n, r, u, M, b, c, h bool

	//Переделать log.Fatal на panic при парсе

	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки") // сравнить соседние
	flag.BoolVar(&M, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&b, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&c, "c", false, "проверять отсортированы ли данные")                    // функция сравнения для поиска нарушения сортировки
	flag.BoolVar(&h, "h", false, "сортировать по числовому значению с учетом суффиксов") //

	flag.IntVar(&k, "k", 0, "указание колонки для сортировки")

	flag.StringVar(&name, "name", "test.txt", "путь к файлу")

	flag.Parse()

	//name := "test.txt"
	data := readFile(name, b)

	fmt.Println(data)
	fmt.Println("")
	fmt.Println(data[4])
	fmt.Println("")
	if n == true {
		sortForNumber(data, k, r)
	}
	if M == true {
		sortForMonth(data, k, r)
	}
	fmt.Println(data)
}

//Построчное чтение из файла
func readFile(name string, b bool) (dataT [][]string) {
	file, err := os.OpenFile(name, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data := make([]string, 0, 4)
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if b == true {
			data = append(data, strings.TrimSpace(line))
		} else {
			data = append(data, line)
		}
		fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
	}
	dataT = dataTable(data)
	return
}

//Делает из слайса строк, слайс слайсов слов
func dataTable(d []string) (t [][]string) {
	t = make([][]string, len(d))
	for a, valA := range d {
		val := strings.Fields(valA)
		t[a] = make([]string, len(val))
		for b, valB := range val {
			t[a][b] = valB
		}
	}
	return
}

//Числовая сортировка, принимает слайс слайсов слов, индекс колонки, порядок сортировки(r == true - обратная сортировка)
func sortForNumber(t [][]string, k int, r bool) {
	sort.Slice(t, func(i, j int) bool {
		a, err := strconv.Atoi(t[i][k])
		if err != nil {
			log.Fatalf("Неверный формат столбца для сортировки: %v\n", err)
		}
		b, err := strconv.Atoi(t[j][k])
		if err != nil {
			log.Fatalf("Неверный формат столбца для сортировки: %v\n", err)
		}

		if r == true {
			return a > b
		} else {
			return a < b
		}
	})
}

//Сортировка по месяцам, принимает слайс слайсов слов, индекс колонки, порядок сортировки(r == true - обратная сортировка)
func sortForMonth(t [][]string, k int, r bool) {
	var month = map[string]int{
		"january":   1,
		"february":  2,
		"mart":      3,
		"april":     4,
		"may":       5,
		"june":      6,
		"july":      7,
		"august":    8,
		"september": 9,
		"october":   10,
		"november":  11,
		"december":  12,
	}
	sort.Slice(t, func(i, j int) bool {
		var a, b int
		a = month[t[i][k]]
		b = month[t[j][k]]

		if r == true {
			return a > b
		} else {
			return a < b
		}
	})
}
