package models

import "net/http"

type Result struct {
	URL      string
	Response *http.Response
	Error    error
}
