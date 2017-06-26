package tweety

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
)

type Tweety struct {
	api *anaconda.TwitterApi
}

type Conf struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func NewClient(conf *Conf) *Tweety {
	anaconda.SetConsumerKey(conf.ConsumerKey)
	anaconda.SetConsumerSecret(conf.ConsumerSecret)
	api := anaconda.NewTwitterApi(conf.AccessToken, conf.AccessTokenSecret)

	log := &logger{logrus.New()}
	api.SetLogger(log)

	return &Tweety{
		api: api,
	}
}
