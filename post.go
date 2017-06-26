package tweety

import (
	"flag"
	"net/url"

	"github.com/yanzay/log"
)

// Post allow us to write a new post on a Twitter account.
func (t *Tweety) Post(args []string) {
	fs := flag.NewFlagSet("post", flag.ExitOnError)
	message := fs.String("message", "", "message to post in twitter.")

	fs.Parse(args)

	if fs.NArg() != 0 || *message == "" {
		fs.Usage()
	}

	log.Infof("\nNew message: %s\n", *message)
	tweet, err := t.api.PostTweet(*message, url.Values{})
	if err != nil {
		log.Errorf("could not tweet '%s': %v", *message, err)
	}

	log.Infof("\nTweet created at:%s\n", tweet.CreatedAt)
}
