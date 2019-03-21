package main

import (
	"flag"
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	log "github.com/sirupsen/logrus"
)

var (
	logger = log.New()
)

func getTimeline(api *anaconda.TwitterApi) ([]anaconda.Tweet, error) {
	args := url.Values{}
	args.Add("count", "3200")       // Twitter only returns most recent 20 tweets by default, so override
	args.Add("include_rts", "true") // When using count argument, RTs are excluded, so include them as recommended
	timeline, err := api.GetUserTimeline(args)
	if err != nil {
		return make([]anaconda.Tweet, 0), err
	}
	return timeline, nil
}

func main() {
	var tweetDays = flag.String("d", "28d", "days after which to delete a tweet")
	flag.Parse()

	anaconda.SetConsumerKey(os.Getenv("TC_TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TC_TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TC_TWITTER_ACCESS_TOKEN"), os.Getenv("TC_TWITTER_ACCESS_TOKEN_SECRET"))
	api.SetLogger(anaconda.BasicLogger)

	fmter := new(log.TextFormatter)
	fmter.FullTimestamp = true
	log.SetFormatter(fmter)
	log.SetLevel(log.InfoLevel)

	ageLimit, err := time.ParseDuration(*tweetDays)
	if err != nil {
		log.Error("Could not parse tweetDays")
	}

	timeline, err := getTimeline(api)
	if err != nil {
		log.Error("Could not get timeline")
	}

	for _, t := range timeline {
		createdTime, err := t.CreatedAtTime()
		if err != nil {
			log.Error("Couldn't parse time ", err)
		} else {
			if time.Since(createdTime) > ageLimit {
				_, err := api.DeleteTweet(t.Id, true)
				log.Info("DELETED: Age - ", time.Since(createdTime).Round(1*time.Minute), " - ", t.Text)
				if err != nil {
					log.Error("Failed to delete! ", err)
				}
			}
		}
	}
}
