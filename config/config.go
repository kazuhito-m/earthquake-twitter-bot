package config

import (
	"github.com/BurntSushi/toml"
)

type Settings struct {
	Twitter TwitterSettings
	Bot BotSetting
}

type TwitterSettings struct {
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
}

type BotSetting struct {
	IntervalSecond int
}

func LoadSettings(filePath string) Settings {
	var settings Settings
	toml.DecodeFile(filePath, &settings)
	return settings
}