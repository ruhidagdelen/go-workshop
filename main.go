package main

import (
	_ "encoding/csv"
	"fmt"
	"go-workshop/helpers"
)

func main() {

	_, data := helpers.ReadFile()
	for index, elem := range data {
		if index > 0 && index < 4 {
			err2, responseObj := helpers.GetRequestData(elem[1])
			if err2 != nil {
				fmt.Println("we got boom errored")
				continue
			}
			fmt.Println(responseObj)
		}
	}

}
