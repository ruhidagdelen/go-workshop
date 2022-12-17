package main

import (
	"fmt"
	"math/rand"
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

	//var helloString string

	//helloString = fmt.Sprintf("%s %s", quote.Hello(), time.Now())

	rand.Seed(time.Now().UnixNano())
	//
	fmt.Println("A dice rolled: ", rollDice())

	fmt.Printf("2 Dices Rolled: %v\n", rollDices(2))
	//
	//fmt.Printf("3 Dices Rolled: %v\n", rollDices(3))
	//
	//fmt.Printf("4 Dices Rolled: %v\n", rollDices(4))

}
