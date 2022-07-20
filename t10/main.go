package main

/*
Простейший telnet клиент
*/

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/reiver/go-telnet"
)

// go run main.go go-telnet --timeout=10s host port go-telnet djxmmx.net 17

func main() {
	var host, port string
	var errTimeout, err error
	var timeout int
	var conn *telnet.Conn

	timeout, errTimeout = strconv.Atoi(os.Args[2][10 : len(os.Args[2])-1])
	if errTimeout != nil {
		log.Println("Неверно задан timeout")
		timeout = 10
	}
	fmt.Println(timeout)

	if os.Args[3] != "host" {
		host = os.Args[3]
		port = os.Args[4]

		fmt.Println(host)
		fmt.Println(port)
		fmt.Println(timeout)

	} else {
		if os.Args[3] == "host" && os.Args[4] == "port" && os.Args[5] == "go-telnet" {
			host = os.Args[6]
			port = os.Args[7]
			fmt.Println(host)
			fmt.Println(port)

		} else {
			os.Exit(1)
		}
	}
	timeIsOver := time.After(time.Duration(timeout * int(time.Second)))
	fmt.Println(time.Duration(timeout * int(time.Second)))

	greacefulShutdown := make(chan os.Signal, 1)

	done := make(chan bool)
	signal.Notify(greacefulShutdown, os.Interrupt, syscall.SIGINT)
	go func() {
		for {
			select {
			case <-timeIsOver:
				log.Println("Время истекло")
				done <- true
				return
			case <-greacefulShutdown:
				sig := <-greacefulShutdown
				log.Println("\nGot signal:", sig)
				done <- true
				return
			default:
				out := make([]string, 0, 4)
				conn, err = telnet.DialTo(os.Args[6] + ":" + os.Args[7])
				defer conn.Close()
				if err == nil {
					buf := bufio.NewReader(conn)
					for {

						line, err := buf.ReadBytes(byte('\n'))

						if err != nil {
							if err == io.EOF {
								break
							} else {
								log.Printf("Ошибка чтения сокета: %v\n", err)
								return
							}
						}
						out = append(out, string(line))

					}
					fmt.Println(strings.Join(out, ""))
					done <- true
					return
				}

			}
		}
	}()
	<-done
}
