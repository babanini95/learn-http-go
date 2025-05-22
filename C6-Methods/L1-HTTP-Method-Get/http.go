package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	users := make([]User, 0)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return users, err
	}

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&users)
	if err != nil {
		return users, err
	}

	return users, nil
}
