package main

import (
	"fmt"
	"math/rand"
	"rsc.io/quote"
	"strings"
	"time"
)

func rollDice() int {
	var dice int
	dice = rand.Intn(6)
	return dice + 1
}

func rollDices(diceCount int) []int {
	var diceResults []int
	for i := 0; i < diceCount; i++ {
		diceResults = append(diceResults, rollDice())
	}
	return diceResults
}

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
	fmt.Println("---------------------")

	//rand.Seed(time.Now().UnixNano())
	//fmt.Println("A dice rolled: ", rollDice())
	//
	//fmt.Printf("2 Dices Rolled: %v\n", rollDices(2))
	//
	//fmt.Printf("3 Dices Rolled: %v\n", rollDices(3))
	//
	//fmt.Printf("4 Dices Rolled: %v\n", rollDices(4))

	// ARRAY
	var myArray [3]string

	myArray = [3]string{"karadeniz", "teknik", "üniversitesi"}

	//myArray = append(myArray, "trabzon")
	for _, word := range myArray {
		fmt.Printf("%v ", word)
	}

	// SLICE

	var mySlice []string
	for _, word := range myArray {
		mySlice = append(mySlice, word)
	}
	mySlice = append(mySlice, "trabzon")

	concatenedText := strings.Join(mySlice, " ")
	fmt.Println(concatenedText)

}
