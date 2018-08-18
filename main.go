package main

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/config"
	"github.com/kazuhito-m/earthquake-twitter-bot/twitter"
	"os"
)

func main() {
	settingFilePath := os.Args[1]
	settings := config.LoadSettings(settingFilePath)
	twitterTool := twitter.CreateTwitter(settings.Twitter)
	twitterTool.Tweet(settingFilePath)
}
