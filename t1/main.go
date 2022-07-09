package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		log.Fatal(err)
	}
	time := time.Now().Add(response.ClockOffset)
	fmt.Println(time)

}
