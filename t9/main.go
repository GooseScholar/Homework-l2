package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
Утилита cut
*/

func main() {
	Wget("https://translate.google.ru/?hl=ru&tab=rT", "wget")
}

const (
	bufSize = 1024 * 8
)

//Все результаты записываются в файл с именем в запросе. Если имя не задано, то используется окончание ссылки
//например: https://translate.google.ru/?hl=ru&tab=rT, название: ?hl=ru&tab=rT
func Wget(url, fileName string) {
	resp := getResponse(url)
	if fileName == "" {
		urlSplit := strings.Split(url, "/")
		fileName = urlSplit[len(urlSplit)-1]
	}
	writeToFile(fileName, resp)
}

// делаем запрос к url и возвращаем ответ
func getResponse(url string) *http.Response {
	tr := new(http.Transport)
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	errorChecker(err)
	return resp
}

// запись ответа в файл
func writeToFile(fileName string, resp *http.Response) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	errorChecker(err)
	defer file.Close()
	bufferedWriter := bufio.NewWriterSize(file, bufSize)
	errorChecker(err)
	_, err = io.Copy(bufferedWriter, resp.Body)
	errorChecker(err)
}

// Проверка ошибок
func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
