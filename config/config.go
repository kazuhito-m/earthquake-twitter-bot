package config

import (
	"github.com/BurntSushi/toml"
)

type Settings struct {
	Twitter TwitterSettings
}

type TwitterSettings struct {
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
}

func LoadSettings(filePath string) Settings {
	var settings Settings
	toml.DecodeFile(filePath, &settings)
	return settings
}