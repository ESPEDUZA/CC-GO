package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

type Repository struct {
	Name        string `json:"name"`
	CloneURL    string `json:"clone_url"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updated_at"`
}

func FetchRepositories(user string, token string) ([]Repository, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s/repos?sort=updated&per_page=100", user), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "token "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var repos []Repository
	err = json.Unmarshal(body, &repos)
	if err != nil {
		return nil, err
	}

	// Trier les répos par date de dernière modification
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].UpdatedAt > repos[j].UpdatedAt
	})

	return repos, nil
}
