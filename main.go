package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
	"github.com/yanzay/log"
)

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}

const (
	appUsage = `
tweety is command line tool.
Usage:
		tweety command [options] [...args]
The commands are:
`
	cmdsHelp = `
Use "tweety command -h" for more information about a command.
The options are:
`
	serveUsage = `
usage: tweety post "message"
post creates a new tweet and send to the main account.
The options are:
`
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, appUsage)
		for _, c := range subcmds {
			fmt.Fprintf(os.Stderr, "    %-24s %s\n", c.name, c.description)
		}

		fmt.Fprintln(os.Stderr, cmdsHelp)
		flag.PrintDefaults()
		os.Exit(1)
	}
}

type subcmd struct {
	name        string
	description string
	run         func(args []string, api *anaconda.TwitterApi)
}

var subcmds = []subcmd{
	{"post", "post a new tweet.", tweetyPost},
	{"sniff", "sniff a word in twitter stream.", tweetySniff},
}

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

func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	log := &logger{logrus.New()}
	api.SetLogger(log)

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}

	subcmd := flag.Arg(0)
	for _, c := range subcmds {
		if c.name == subcmd {
			c.run(flag.Args()[1:], api)
			return
		}
	}

	fmt.Fprintf(os.Stderr, "unknown subcmd %q\n", subcmd)
	fmt.Fprintln(os.Stderr, `Run "tweety -h" for usage.`)
	os.Exit(1)
}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }
