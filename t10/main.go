package main

/*
Утилита telnet
*/

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/reiver/go-telnet"
)

func main() {
	var timeout, host, port string
	timeout = os.Args[2]
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
			fmt.Println(timeout[10 : len(timeout)-1])
		} else {
			os.Exit(1)
		}
	}
	fmt.Println(os.Args[6] + " " + os.Args[7])

	//var caller telnet.Caller = telnet.StandardCaller
	//b := make([]byte, 0)
	conn, _ := telnet.DialTo(os.Args[6] + ":" + os.Args[7])

	//var err error
	//var n int
	//
	//for nil == err {
	//	n, err = conn.Read(b)
	//	fmt.Println(string(b[:n]))
	//}
	io.Copy(os.Stdout, conn)

}

func copyTo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

/*
var handler telnet.Handler = telnet.EchoHandler

	err := telnet.ListenAndServe(":5555", handler)
	if nil != err {
		//@TODO: Handle this error better.
		panic(err)
	}
*/

//52.88.68.92:1234

//159.69.204.41:22
//"opennet.ru:80"
