package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id

	userData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(userData))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, nil
	}
	defer resp.Body.Close()

	var user User
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&user)

	return user, err
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, nil
	}

	req.Header.Set("X-API-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, nil
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var user User
	err = decoder.Decode(&user)
	return user, err
}
