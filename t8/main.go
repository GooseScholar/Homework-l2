package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	ps "github.com/mitchellh/go-ps"
	process "github.com/shirou/gopsutil/v3/process"
)

func main() {
	quit := make(chan bool)
	//scanner := bufio.NewScanner(os.Stdin)
	r := bufio.NewReader(os.Stdin)
	go func() {

		fmt.Println("Введите команду:")
		for {
			in, _ := r.ReadString('\n')

			command := strings.Fields(in)

			switch command[0] {
			case "quite":
				quit <- true
				return
			case "pwd":
				fmt.Println(getWd())
			case "echo":

				if len(command) != 1 {
					fmt.Println(command[1:])

				} else {
					fmt.Fprintln(os.Stderr, "invalid entry: expect more than 1 value, have 1")
				}
			case "cd":
				if len(command) == 2 {
					chdir(command[1])
				} else {
					fmt.Fprintf(os.Stderr, "invalid entry: expect 2 values, have "+"%d", len(command))
				}
			case "kill":
				if len(command) != 1 {
					err := kill(string(command[1]))
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
					}
				}
			case "ps":
				psGo()
			case "ls":
				ls(getWd())
			}

		}
	}()
	//..
	<-quit
}

//Вывести текущую директорию
func getWd() string {
	dir, err := os.Getwd()
	if err != nil {
		println(err)
	}
	return dir
}

//Вывести в stdout
func echo(text ...string) {
	var textBytes []byte
	for i, t := range text {
		if i > 0 {
			textBytes = append(textBytes, ' ')
		}
		textBytes = append(textBytes, t...)
	}
	textBytes = append(textBytes, '\n')
	if _, err := os.Stdout.Write(textBytes); err != nil {
		panic(err)
	}
}

//Показать все файлы текущей директории
func ls(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	list, err := file.Readdirnames(0)
	if err != nil {
		panic(err)
	}
	sort.Strings(list)
	for i := len(list) - 1; i >= 0; i-- {
		name := list[i]
		fmt.Println(name)
	}
}

//Изменить текущую директорию
func chdir(cd string) {
	os.Chdir(cd)
	newDir, err := os.Getwd()
	if err != nil {
	}
	fmt.Println(newDir)
}

//Вывести список процессов
func psGo() {
	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed")
		return
	}
	for x := range processList {
		var process ps.Process
		process = processList[x]
		log.Printf("%d\t%s\n", process.Pid(), process.Executable())

	}
}

//Убить процесс
func kill(name string) error {
	processes, err := process.Processes()
	if err != nil {
		return err
	}
	for _, p := range processes {
		n, err := p.Name()
		if err != nil {
			return err
		}
		if n == name {
			return p.Kill()
		}
	}
	return fmt.Errorf(" process not found")
}
