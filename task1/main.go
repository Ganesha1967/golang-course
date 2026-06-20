package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type GithubRepository struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Stars        int    `json:"stargazers_count"`
	Forks        int    `json:"forks_count"`
	CreationDate string `json:"created_at"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <github_repo_url>\n", os.Args[0])
		return
	}

	repoURL := os.Args[1]
	owner, repo, err := parseGitHubURL(repoURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	repository, err := parseRepoData(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	printInfo(repository)
}

func parseGitHubURL(repoURL string) (owner, repo string, err error) {
	if !strings.Contains(repoURL, "https://") {
		return "", "", fmt.Errorf("URL must start with https://")
	}

	splitedURL := strings.Split(repoURL, "/")
	if len(splitedURL) < 5 {
		return "", "", fmt.Errorf("invalid URL")
	}
	if splitedURL[2] != "github.com" {
		return "", "", fmt.Errorf("URL must be from GitHub")
	}

	owner, repo = splitedURL[3], splitedURL[4]
	if owner == "" || repo == "" {
		return "", "", fmt.Errorf("invalid owner or repository name")
	}

	return owner, repo, nil
}

func parseRepoData(url string) (*GithubRepository, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("repository not found")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected API response: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var repository GithubRepository
	err = json.Unmarshal(data, &repository)
	if err != nil {
		return nil, fmt.Errorf("invalid JSON response: %v", err)
	}

	return &repository, nil
}

func printInfo(repo *GithubRepository) {
	fmt.Printf(`
Name:         %s
Description:  %s
Stars:        %d
Forks:        %d
Created:      %s
`, repo.Name, repo.Description, repo.Stars, repo.Forks, repo.CreationDate,
	)
}
