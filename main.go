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
		log.Println(*repo.Name)
	}
}
