package main

import (
	"context"
	"github.com/google/go-github/github"
	"log"
)

func main() {
	ctx := context.Background()
	client := github.NewClient(nil)
	username := "randomecho"

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
