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
	if !validate() {
		return
	}
	settingFilePath := os.Args[1]
	settings := config.LoadSettings(settingFilePath)
	twitter := twitter.CreateTwitter(settings.Twitter)
	site := earthquake.CreateSite(web.CreateClient())
	bot := bot.CreateBot(settings, twitter, site)

	bot.Run()
}

func validate() bool {
	if len(os.Args) < 2 {
		fmt.Println("引数がありません。引数に設定ファイルのPathを指定して下さい。")
		return false
	}
	return true
}
