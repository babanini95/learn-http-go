package main

func fetchTasks(baseURL, availability string) []Issue {
	baseURL += "?sort=estimate"
	switch availability {
	case "Low":
		baseURL += "&limit=1"
	case "Medium":
		baseURL += "&limit=3"
	case "High":
		baseURL += "&limit=5"
	}

	fullURL := baseURL
	return getIssues(fullURL)
}
