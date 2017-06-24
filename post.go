package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/yanzay/log"
)

// tweetyPost this method allow us to write a new post on a Twitter account.
func tweetyPost(args []string, api *anaconda.TwitterApi) {
	fs := flag.NewFlagSet("post", flag.ExitOnError)
	message := fs.String("message", "", "message to post in twitter.")

	fs.Parse(args)

	if fs.NArg() != 0 || *message == "" {
		fs.Usage()
	}
	fmt.Printf("\nNew message: %s\n", *message)
	tweet, err := api.PostTweet(*message, url.Values{})
	if err != nil {
		log.Errorf("could not tweet '%s': %v", *message, err)
	}

	fmt.Printf("\nTweet created at:%s\n", tweet.CreatedAt)
}
