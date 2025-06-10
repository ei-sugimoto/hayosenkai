package main

import (
	"flag"
	"log"

	"github.com/slack-go/slack"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("No arguments provided")
	}

	slackToken := args[0]

	s := slack.New(slackToken)

}
