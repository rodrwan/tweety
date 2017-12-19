package tweety

import (
	"flag"
	"fmt"
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
		return
	}

	newMsg := fmt.Sprintf("%s #bytweety", *message)
	log.Infof("\nNew message: %s\n", newMsg)
	tweet, err := t.api.PostTweet(newMsg, url.Values{})
	if err != nil {
		log.Errorf("could not tweet '%s': %v", newMsg, err)
	}

	log.Infof("Tweet created at:%s\n", tweet.CreatedAt)
}
