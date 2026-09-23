package ui

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func fetchREADME(repoURL string) (string, error) {
	repoURL = strings.TrimSuffix(repoURL, "/")

	parts := strings.Split(strings.TrimPrefix(repoURL, "https://github.com/"), "/")

	if len(parts) < 2 {
		return "", fmt.Errorf("invalid URL")
	}

	owner := parts[0]
	repo := parts[1]

	url := fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/%s/HEAD/README.md",
		owner,
		repo,
	)

	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("README: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("README ERROR: %w", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("READING README: %w", err)
	}

	return string(body), nil
}

func updateREADMEs() {
	for i := range projectItems {
		if projectItems[i].RepoURL == "" {
			continue
		}

		readme, err := fetchREADME(projectItems[i].RepoURL)
		if err != nil {
			continue
		}

		if readme != projectItems[i].README {
			projectItems[i].README = readme
		}
	}
}

func PollREADME(interval time.Duration) {
	updateREADMEs()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		updateREADMEs()
	}
}
