package helpers

import (
	"fmt"
	"github.com/asmcos/requests"
	"go-workshop/models"
)

func GetRequestData(domain string) (error, models.WebsiteResponse) {
	var importantResponse models.WebsiteResponse
	requestsItem := requests.Requests()

	url := fmt.Sprintf("https://%s", domain)
	response, err := requestsItem.Get(url)
	if err != nil {
		fmt.Println("Request get error", err)
		return err, importantResponse
	}
	//importantResponse := models.WebsiteResponse{}
	importantResponse.Text = response.Text()
	importantResponse.ResponseCode = response.R.StatusCode

	return nil, importantResponse
}
