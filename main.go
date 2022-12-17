package main

import "fmt"
import "rsc.io/quote"
import "time"

func main() {
	fmt.Println("Hello World")

	//var helloString string
	//
	//helloString = fmt.Sprintf("%s %s", quote.Hello(), time.Now())

	var helloString string

	fillerText := "The exact time is"

	helloString = fmt.Sprintf("%s %s %s", quote.Hello(), fillerText, time.Now())

	fmt.Println(helloString)

	theBestInteger := 61.0

	//theBestInteger = 61.5

	theSecondBest := 61.5

	sumOfThem := theBestInteger + theSecondBest

	fmt.Println(sumOfThem)

	var a string
	a = ""
	_ = a

	b := 1.2
	_ = b
}
