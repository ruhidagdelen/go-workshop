package models

type WebsiteResponseTable struct {
	Id           int    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Text         string `gorm:"type:VARCHAR;column:text" json:"text"`
	ResponseCode int    `gorm:"type:INT4;column:response_code" json:"response_code"`
}
