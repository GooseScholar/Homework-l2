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
	signal.Notify(greacefulShutdown, syscall.SIGQUIT)
	go func() {
		for {
			select {
			case <-timeIsOver:
				fmt.Println("Время истекло")
				done <- true
				return
			case <-greacefulShutdown:
				fmt.Println("Сигнал остановки")
				done <- true
				return
			default:
				out := make([]string, 0, 4)
				conn, err = telnet.DialTo(os.Args[6] + ":" + os.Args[7])
				defer conn.Close()
				if err == nil {
					buf := bufio.NewReader(conn)
					for {
						fmt.Println("1")
						line, err := buf.ReadBytes(byte('\n'))
						fmt.Println("2")
						if err != nil {
							if err == io.EOF {
								break
							} else {
								log.Printf("Ошибка чтения сокета: %v\n", err)
								return
							}
						}
						out = append(out, string(line))
						//time.Sleep(500 * time.Millisecond)
					}
					fmt.Println(strings.Join(out, ""))
					done <- true
					return
				}

			}
		}
	}()
	<-done

	//conn, _ := telnet.DialTo(os.Args[6] + ":" + os.Args[7])
	//conn, _ := telnet.DialTimeout(os.Args[6]+":"+os.Args[7], time.Duration(timeout*int(time.Second)))
	//defer conn.Close()

}

//52.88.68.92:1234
//djxmmx.net:17
//159.69.204.41:22
//"opennet.ru:80"

func DoWithTries(fn func() error, delay time.Duration) (err error) {
	for {
		if err = fn(); err != nil {
			time.Sleep(500 * time.Millisecond)

			continue
		}
		return
	}

}

/*
err = DoWithTries(func() error {
		_, cancle := context.WithTimeout(context.Background(), time.Duration(timeout*int(time.Second)))
		defer cancle()

		conn, err = telnet.DialTo(os.Args[6] + ":" + os.Args[7])
		if err != nil {
			return err
		}
		defer conn.Close()
		buf := bufio.NewReader(conn)
		for {
			line, err := buf.ReadBytes(byte('\n'))

			out = append(out, string(line))

			if err != nil {
				if err == io.EOF {
					break
				} else {
					log.Printf("Ошибка чтения сокета: %v\n", err)
					return err
				}
			}

		}
		fmt.Println(strings.Join(out, ""))
		return nil
	}, time.Duration(timeout*int(time.Second)))
	if err != nil {
		log.Println(err)
		conn.Close()
		return
	}
*/
/*
out := make([]string, 0, 4)
	conn, err = telnet.DialTo(os.Args[6] + ":" + os.Args[7])
	if err != nil {
		return
	}
	defer conn.Close()
	buf := bufio.NewReader(conn)
	for {
		line, err := buf.ReadBytes(byte('\n'))

		out = append(out, string(line))

		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Printf("Ошибка чтения сокета: %v\n", err)
				return
			}
		}

	}
	fmt.Println(strings.Join(out, ""))
*/
