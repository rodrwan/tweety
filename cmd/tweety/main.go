package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rodrwan/tweety"
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
	run         func(args []string)
}

var subcmds = []subcmd{}

// main program
func main() {
	tConf := &tweety.Conf{
		ConsumerKey:       consumerKey,
		ConsumerSecret:    consumerSecret,
		AccessToken:       accessToken,
		AccessTokenSecret: accessTokenSecret,
	}

	t := tweety.NewClient(tConf)

	subcmds = []subcmd{
		{"post", "post a new tweet.", t.Post},
		{"sniff", "sniff a word in twitter stream.", t.Sniff},
		{"search", "search tweets that contain a certain word", t.Search},
		{"rt", "retweet a tweet by it ID", t.Retweet},
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}

	subcmd := flag.Arg(0)
	for _, c := range subcmds {
		if c.name == subcmd {
			c.run(flag.Args()[1:])
			return
		}
	}

	fmt.Fprintf(os.Stderr, "unknown subcmd %q\n", subcmd)
	fmt.Fprintln(os.Stderr, `Run "tweety -h" for usage.`)
	os.Exit(1)
}
