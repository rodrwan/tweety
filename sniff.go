package tweety

import (
	"flag"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/yanzay/log"
)

// Sniff listen tweets that contain the given word.
// Inspired by #JustForFunc of Francesc Campoy (https://www.youtube.com/c/justforfunc)
func (t *Tweety) Sniff(args []string) {
	fs := flag.NewFlagSet("sniff", flag.ExitOnError)
	word := fs.String("word", "", "word to sniff in twitter stream.")

	fs.Parse(args)

	if fs.NArg() != 0 || *word == "" {
		fs.Usage()
	}
	log.Infof("\nListening to messages containing: %s\n", *word)
	stream := t.api.PublicStreamFilter(url.Values{
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

		log.Infof("\nUser: %s | ID: %d\n%s\n\n", t.User.Name, t.Id, t.Text)
	}
}
