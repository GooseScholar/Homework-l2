package main

import "github.com/reiver/go-telnet"

func main() {
	var caller telnet.Caller = telnet.StandardCaller

	//@TOOD: replace "example.net:5555" with address you want to connect to.
	telnet.DialToAndCall("52.88.68.92:1234", caller)
}

//52.88.68.92:1234
//159.69.204.41:22
//"opennet.ru:80"

/*
connbuf := bufio.NewReader(c.m_socket)
	// Read the first byte and set the underlying buffer
	b, _ := connbuf.ReadByte()
	if connbuf.Buffered() > 0 {
		var msgData []byte
		msgData = append(msgData, b)
		for connbuf.Buffered() > 0 {
			// read byte by byte until the buffered data is not empty
			b, err := connbuf.ReadByte()
			if err == nil {
				msgData = append(msgData, b)
			} else {
				log.Println("-------> unreadable caracter...", b)
			}
		}
		// msgData now contain the buffered data...

	}
*/
