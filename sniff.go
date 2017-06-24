package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/yanzay/log"
)

// tweetySniff this method listen tweets that contain the given word.
// Inspired by #JustForFunc of Francesc Campoy (https://www.youtube.com/c/justforfunc)
func tweetySniff(args []string, api *anaconda.TwitterApi) {
	fs := flag.NewFlagSet("sniff", flag.ExitOnError)
	word := fs.String("word", "", "word to sniff in twitter stream.")

	fs.Parse(args)

	if fs.NArg() != 0 || *word == "" {
		fs.Usage()
	}
	fmt.Printf("\nListening to messages containing: %s\n", *word)
	stream := api.PublicStreamFilter(url.Values{
		"track": []string{*word},
	})

	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			log.Warningf("received unexpected value of type %T", v)
			continue
		}

		if t.RetweetedStatus != nil {
			continue
		}

		log.Infof("\nUser: %s\n%s\n\n", t.User.Name, t.Text)
	}
}
