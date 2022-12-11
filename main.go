package main

import "fmt"

func brainwash(saying *string) {
	*saying = "Beep Boop."
}

func main() {
	greeting := "Hello there!"

	brainwash(&greeting)

	fmt.Println("greeting is now:", greeting)
}
