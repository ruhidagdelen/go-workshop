package main

import (
	_ "encoding/csv"
	"fmt"
	"go-workshop/helpers"
	"go-workshop/models"
)

func main() {
	helpers.InitializeDB()

	db := helpers.GetDB()

	_ = db.AutoMigrate(&models.WebsiteResponseTable{})

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
