package main

import "fmt"
import "rsc.io/quote"
import "time"

func main() {
	fmt.Println("Hello World")

	fmt.Println(quote.Go())

	fmt.Println(quote.Hello())

	fmt.Println(time.Now())
}
