package config

import (
	"github.com/BurntSushi/toml"
)

type Settings struct {
	Twitter TwitterSettings
	Bot BotSetting
	GoogleApi GoogleApiSettings
}

type TwitterSettings struct {
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
}

type GoogleApiSettings struct {
	ApiKey string
}

type BotSetting struct {
	IntervalSecond int
	FrequentThresholdHour int
}

func LoadSettings(filePath string) Settings {
	var settings Settings
	toml.DecodeFile(filePath, &settings)
	return settings
}