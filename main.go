package main

import (
	"fmt"
	"math/rand"
	"strings"
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

	var sampleInt int
	var sampleString string
	var sampleFloat float32

	fmt.Println("default values")
	fmt.Printf("default int '%v'\n", sampleInt)
	fmt.Printf("default string '%v'\n", sampleString)
	fmt.Printf("default float '%v'\n", sampleFloat)

	//ARRAY
	var myArray [3]string

	myArray = [3]string{"karadeniz", "üniversitesi"}

	//myArray = append(myArray, "trabzon")
	for _, word := range myArray {
		fmt.Printf("%v ", word)
	}
	//
	//// SLICE
	//
	var mySlice []string
	for _, word := range myArray {
		mySlice = append(mySlice, word)
	}
	mySlice = append(mySlice, "trabzon")
	//
	concatenedText := strings.Join(mySlice, " ")
	fmt.Println(concatenedText)

}
