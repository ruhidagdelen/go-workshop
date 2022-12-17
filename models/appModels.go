package models

import "net/http"

type WebsiteResponse struct {
	Text         string
	ResponseCode int
	Headers      http.Header
	Cookie       http.Cookie
}
