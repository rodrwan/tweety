package tweety

import (
	"flag"

	"github.com/yanzay/log"
)

// Retweet allow to retweet a tweet by it id.
func (t *Tweety) Retweet(args []string) {
	fs := flag.NewFlagSet("rt", flag.ExitOnError)
	id := fs.Int64("id", 0, "tweet id to retweet.")

	fs.Parse(args)

	if fs.NArg() != 0 || *id == 0 {
		fs.Usage()
	}

	log.Infof("\nRetweet: %d\n", *id)
	tt, err := t.api.Retweet(*id, false)
	if err != nil {
		log.Errorf("could not retweet %d: %v", *id, err)
		return
	}
	log.Infof("retweeted %d\nMessage: %s", *id, tt.Text)
}
