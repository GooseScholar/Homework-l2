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

func (p pkgString) Unpack() string {
	//Если ничего, то ничего
	if p == "" {
		return ""
	}
	//Строку на входе превращаем в слайс рун
	message := []rune(p)
	//Слайс рун в который записываем процесс распаковывания
	result := make([]rune, 0, 25)
	//Отслеживет появление `\`
	var esc bool

	for i, oneRune := range message {
		//если первый символ строки это число, то выводим ошибку
		if unicode.IsDigit(oneRune) && i == 0 {
			return "(некорректная строка)"
		}

		//если пришла буква, то пишем её
		if unicode.IsLetter(oneRune) {
			result = append(result, oneRune)

		}
		if unicode.IsDigit(oneRune) {
			//распознаем цифру
			counter, err := strconv.Atoi(string(oneRune))
			//возвращение ошибки, если не удалось конвертировать руну в int
			if err != nil {
				log.Printf("Ошибка конвертации в int: %v", err)
				return "(некорректная строка)"
			}

			//Если предыдущий символ `\`
			if esc == true {
				result = append(result, oneRune)
				esc = false
			} else if counter > 1 {
				//дописать оставшиеся буквы (число минус 1, т.к. предыдущий сымвол уже записан)
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
