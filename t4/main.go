package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Поиск анаграмм по словарю
*/

func main() {
	slice := []string{"тяпка", "листок", "листик", "Столик", "Слиток", "вертикаль", "материк", "пятка", "Керамит", "Мошкара", "Ромашка", "Торнадо",
		"спаниель", "норматив", "хористка", "минотавр"}
	m := AnagramsSearch(slice)
	fmt.Println(m)
}

//AnagramsSearch поиск и возврат всех анаграмм по группам
func AnagramsSearch(slice []string) (answer map[string][]string) {
	answer = make(map[string][]string)

	m := make(map[string]map[string]struct{})

	for _, word := range slice { //первый ключ - отсортированные слова, второй ключ - само слово.
		lowerWord := strings.ToLower(word)       //сделать все символы маленькими
		sortWord := strings.Split(lowerWord, "") //из строки в слайс строк
		sort.Strings(sortWord)                   // сортировка символов
		_, ok := m[strings.Join(sortWord, "")]
		if ok != true { //если ключа ещё нет, то инициализировать мапу для второго ключа
			m[strings.Join(sortWord, "")] = make(map[string]struct{})
		}

		m[strings.Join(sortWord, "")][lowerWord] = struct{}{}
	}

	var slice2 []string
	for _, keys := range m { //в цикле по первому ключу, запись вторых ключей с слайс и сортировка, после чего запись в мап
		if len(keys) > 1 {
			slice2 = make([]string, 0, len(keys)) //слайс для сортировки множеств
			for k := range keys {
				slice2 = append(slice2, k)
			}
			sort.Strings(slice2)          //сортировка множеств
			for _, word := range slice2 { //запись отсортированных множеств в карту
				answer[slice2[0]] = append(answer[slice2[0]], word)
			}
		}
	}

	return

}
