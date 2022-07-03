package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

type pkgString string

func main() {

	fmt.Println("\nЗадача на распаковку \nДля выхода нажмите CTRL-C")
	fmt.Println("")

	var message pkgString

	for {
		fmt.Print("Введите корректную строку: ")

		//запись в консоль и в переменную
		_, err := fmt.Scanf("%s", &message)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Распакованная строка: ", message.Unpack())
			fmt.Println("")
		}
	}
}

//Функция распаковки
func (p pkgString) Unpack() string {
	if p == "" { //Если ничего, то ничего
		return ""
	}

	message := []rune(p)          //Строку на входе превращаем в слайс рун
	result := make([]rune, 0, 25) //Слайс рун в который записываем процесс распаковывания
	var esc bool                  //Отслеживет появление `\`

	for i, oneRune := range message {
		if unicode.IsDigit(oneRune) && i == 0 { //если первый символ строки это число, то выводим ошибку
			return "(некорректная строка)"
		}
		if unicode.IsLetter(oneRune) { //если пришла буква, то пишем её
			result = append(result, oneRune)

		}
		if unicode.IsDigit(oneRune) {
			counter, err := strconv.Atoi(string(oneRune)) //распознаем цифру
			if err != nil {                               //возвращение ошибки, если не удалось конвертировать руну в int
				log.Printf("Ошибка конвертации в int: %v", err)
				return "(некорректная строка)"
			}
			if esc == true { //Если предыдущий символ `\`
				result = append(result, oneRune)
				esc = false
			} else if counter > 1 { //дописать оставшиеся буквы (число минус 1, т.к. предыдущий сымвол уже записан)
				for j := 0; j < counter-1; j++ {
					result = append(result, message[i-1])
				}
			} else if counter == 0 { //если пришел 0
				result = result[:len(result)-1]
			}

		}
		if oneRune == '\\' { // если пришел символ `\`
			if message[i-1] == '\\' { // если пришел второй символ `\` подряд
				result = append(result, message[i-1])
				esc = false
			} else { // если пришел символ `\`, а предыдущий отличный от него, либо не существует
				esc = true
			}
		}
	}
	return string(result)
}
