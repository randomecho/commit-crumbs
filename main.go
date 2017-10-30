package main

import (
	"context"
	"encoding/json"
	"github.com/google/go-github/github"
	"log"
	"os"
)

type Config struct {
	Username string `json:username`
}

func getConfig(file string) Config {
	var config Config

	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	jsonParse := json.NewDecoder(configFile)
	jsonParse.Decode(&config)

	return config
}

func main() {
	config := getConfig("./config.json")
	ctx := context.Background()
	client := github.NewClient(nil)
	username := config.Username

	opt := &github.RepositoryListOptions{Type: "public"}
	repos, _, err := client.Repositories.List(ctx, username, opt)

	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repos {
		opt := &github.CommitsListOptions{}
		commits, _, err := client.Repositories.ListCommits(ctx, username, *repo.Name, opt)

		if err != nil {
			log.Fatal(err)
		}

		for _, commit := range commits {
			log.Println(*commit.SHA)
			log.Println(*commit.Commit.Author.Name)
			log.Println(*commit.Commit.Message)
		}
	}
}
