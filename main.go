package main

import "fmt"
import "rsc.io/quote"
import "time"

func main() {
	fmt.Println("Hello World")

	//fmt.Println(quote.Go())
	//
	//fmt.Println(quote.Hello())
	//
	//fmt.Println(time.Now())

	var helloString string

	helloString = fmt.Sprintf("%s %s", quote.Hello(), time.Now())

	//var helloString string;
	//
	//fillerText := "The exact time is"
	//
	//helloString = fmt.Sprintf("%s %s %s", quote.Hello(), fillerText, time.Now())

	fmt.Println(helloString)

}
