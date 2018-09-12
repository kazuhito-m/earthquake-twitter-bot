package bot

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/config"
	"github.com/kazuhito-m/earthquake-twitter-bot/earthquake"
	"github.com/kazuhito-m/earthquake-twitter-bot/twitter"
	"time"
)

type PoolingAndAlertBot struct {
	settings       config.Settings
	twitter        twitter.Twitter
	earthquakeSite earthquake.Site
}

func (b PoolingAndAlertBot) Run() {
	intervalSecond := b.settings.Bot.IntervalSecond
	site := b.earthquakeSite
	lastReport := earthquake.EarthquakeReport{}
	messageMaker := CreateMessageMaker(b.settings.Bot)
	for {
		report := site.GetNowEarthquakeAnalyze()
		if !report.Exists() {
			continue
		}
		if lastReport.Same(report) {
			continue
		}

		tweetMessage := messageMaker.Message(report)
		b.twitter.Tweet(tweetMessage)

		lastReport = report
		sleepSecond(intervalSecond)
	}
}

func sleepSecond(second int) {
	time.Sleep(time.Duration(second) * time.Second)
}


func CreateBot(settings config.Settings, twitter twitter.Twitter, earthquakeSite earthquake.Site) PoolingAndAlertBot {
	return PoolingAndAlertBot{settings, twitter, earthquakeSite}
}

