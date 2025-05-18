package main

import (
	"net/http"
)

/*
*
Complete the getContentType function. It takes a pointer to a http.Response as input and should return the Content-Type header.

Use the .Get() method on the Response struct's Header field to get what you need.
*/
func getContentType(res *http.Response) string {
	return res.Header.Get("Content-Type")
}
