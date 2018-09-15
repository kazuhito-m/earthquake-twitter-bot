package main

import (
	"fmt"
	"github.com/kazuhito-m/earthquake-twitter-bot/bot"
	"github.com/kazuhito-m/earthquake-twitter-bot/config"
	"github.com/kazuhito-m/earthquake-twitter-bot/earthquake"
	"github.com/kazuhito-m/earthquake-twitter-bot/twitter"
	"github.com/kazuhito-m/earthquake-twitter-bot/web"
	"os"
)

func main() {
	settingFilePath := os.Args[1]
	settings := config.LoadSettings(settingFilePath)
	twitter := twitter.CreateTwitter(settings.Twitter)
	site := earthquake.CreateSite(web.CreateClient())
	bot := bot.CreateBot(settings, twitter, site)

	bot.Run()
}

