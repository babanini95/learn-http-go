package main

import (
	"net/http"
)

func getUserCode(url string) int {
	res, err := http.Get(url)
	if err != nil {
		return 404
	}
	defer res.Body.Close()

	return res.StatusCode
}
