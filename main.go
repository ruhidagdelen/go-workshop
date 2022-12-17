package main

import (
	"fmt"
	"math/rand"
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
	//fmt.Println("Hello World")

	//fmt.Println(quote.Go())
	//
	//fmt.Println(quote.Hello())
	//
	//fmt.Println(time.Now())

	//var helloString string
	//
	//helloString = fmt.Sprintf("%s %s", quote.Hello(), time.Now())

	//var helloString string;
	//
	//fillerText := "The exact time is"
	//
	//helloString = fmt.Sprintf("%s %s %s", quote.Hello(), fillerText, time.Now())

	//fmt.Println(helloString)
	//fmt.Println("---------------------")

	//rand.Seed(time.Now().UnixNano())
	//fmt.Println("A dice rolled: ", rollDice())
	//
	//fmt.Printf("2 Dices Rolled: %v\n", rollDices(2))
	//
	//fmt.Printf("3 Dices Rolled: %v\n", rollDices(3))
	//
	//fmt.Printf("4 Dices Rolled: %v\n", rollDices(4))

	//var sampleInt int
	//var sampleString string
	//var sampleFloat float32
	//
	//fmt.Println("default values")
	//fmt.Printf("default int '%v'\n", sampleInt)
	//fmt.Printf("default string '%v'\n", sampleString)
	//fmt.Printf("default float '%v'\n", sampleFloat)

	// interface

	var sampleInterface interface{}

	sampleInterface = 5

	fmt.Printf("interface value '%v'\n", sampleInterface)

	sampleInterface = "this is its text"

	fmt.Printf("interface value '%v'\n", sampleInterface)

	// MAPS
	//map string interface
	//var sampleMap map[string]interface{}
	//
	//sampleMap = map[string]interface{}{"integer": 1, "string": "two", "float": 8.2,
	//	"array": [2]string{"one", "two"},
	//	"map":   map[string]interface{}{"literally": "everything"}}
	//
	//for k, v := range sampleMap {
	//	fmt.Printf("key[%s] value[%v]\n", k, v)
	//}
	//
	//// map string string
	//stringMap := map[string]string{"key": "value", "anotherOne": "anotherValue"}
	//
	//for k, v := range stringMap {
	//	fmt.Printf("key[%s] value[%v]\n", k, v)
	//}
	//
	//// map int int
	intMap := map[int]int{0: 1000, 1: 1002}

	for k, v := range intMap {
		fmt.Printf("key[%s] value[%v]\n", k, v)
	}

}
