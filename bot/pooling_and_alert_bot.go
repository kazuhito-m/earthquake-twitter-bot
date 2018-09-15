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
	lastReport := earthquake.EarthquakeReport{}
	messageMaker := CreateMessageMaker(b.settings.Bot)
	for {
		lastReport = checkNewReport(b, lastReport, messageMaker)
		sleepSecond(intervalSecond)
	}
}

func checkNewReport(bot PoolingAndAlertBot, lastReport earthquake.EarthquakeReport, messageMaker MessageMaker) earthquake.EarthquakeReport {
	nowReport := bot.earthquakeSite.GetNowEarthquakeAnalyze()
	if !nowReport.Exists() || lastReport.Same(nowReport) {
		return lastReport
	}

	tweetMessage := messageMaker.Message(nowReport)
	bot.twitter.Tweet(tweetMessage)

	return nowReport
}

func sleepSecond(second int) {
	time.Sleep(time.Duration(second) * time.Second)
}

func CreateBot(settings config.Settings, twitter twitter.Twitter, earthquakeSite earthquake.Site) PoolingAndAlertBot {
	return PoolingAndAlertBot{settings, twitter, earthquakeSite}
}
