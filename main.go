package main

import (
	"fmt"
	"github.com/asmcos/requests"
	"go-workshop/models"
)

func main() {
	fmt.Println("Hello World")

	requestsItem := requests.Requests()

	response, err := requestsItem.Get("https://google.com")
	if err != nil {
		fmt.Println("Request get error", err)
	}

	importantResponse := models.WebsiteResponse{}
	importantResponse.Text = response.Text()
	importantResponse.ResponseCode = response.R.StatusCode
	importantResponse.Headers = response.R.Header

}
