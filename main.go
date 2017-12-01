package main

import (
	"context"
	"encoding/json"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"math/rand"
	"os"
	"time"
)

type Config struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
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
	username := config.Username

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := &github.RepositoryListOptions{Type: "public"}
	repos, _, err := client.Repositories.List(ctx, username, opt)

	if err != nil {
		log.Fatal(err)
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	n := r.Intn(len(repos))
	repo := repos[n]

	optRepo := &github.CommitsListOptions{}
	commits, _, err := client.Repositories.ListCommits(ctx, username, *repo.Name, optRepo)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(*repo.FullName)

	m := make(map[string]map[string]string)

	for _, commit := range commits {
		if m[*commit.Commit.Author.Name] == nil {
			m[*commit.Commit.Author.Name] = map[string]string{}
		}

		m[*commit.Commit.Author.Name][*commit.SHA] = *commit.Commit.Message
	}

	for author, _ := range m {
		log.Println(author, "\n-", len(m[author]), "commits")
	}
}
