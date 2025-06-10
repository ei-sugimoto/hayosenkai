package main

import (
	"context"
	"flag"
	"log"

	"github.com/google/go-github/v72/github"
	"github.com/slack-go/slack"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("No arguments provided")
	}

	slackToken := args[0]
	githubToken := args[1]

	s := slack.New(slackToken)
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
	tc := oauth2.NewClient(ctx, ts)
	gh := github.NewClient(tc)

}
