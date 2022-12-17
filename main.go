package main

import (
	"fmt"
	strftime "github.com/itchyny/timefmt-go"
	"rsc.io/quote"
	"time"
)

func main() {
	fmt.Println("Hello World")

	fmt.Println(quote.Go())

	fmt.Println(quote.Hello())

	fmt.Println(strftime.Format(time.Now(), "%Y-%m-%d %H:%M:%S"))
}
