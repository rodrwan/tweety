package tweety

import (
	"flag"

	"github.com/yanzay/log"
)

// Search search tweets that contain the given word.
func (t *Tweety) Search(args []string) {
	fs := flag.NewFlagSet("search", flag.ExitOnError)
	word := fs.String("word", "", "word to search tweets.")

	fs.Parse(args)

	if fs.NArg() != 0 || *word == "" {
		fs.Usage()
		return
	}
	log.Infof("\nSearching messages containing: %s\n", *word)
	result, err := t.api.GetSearch(*word, nil)
	if err != nil {
		log.Warningf("Something went wrong: %v\n", err)
		return
	}

	for _, tweet := range result.Statuses {
		log.Infof("\nUser: %s\n%s\n", tweet.User.Name, tweet.Text)
	}
}
