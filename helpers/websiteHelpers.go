package helpers

import "go-workshop/models"

func SaveWebsiteResponse(websiteRes models.WebsiteResponse) error {
	var websiteResDb models.WebsiteResponseTable
	db = GetDB()
	websiteResDb.ResponseCode = websiteRes.ResponseCode
	websiteResDb.Text = websiteRes.Text
	db.Save(&websiteResDb)

	return nil
}
