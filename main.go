package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
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
