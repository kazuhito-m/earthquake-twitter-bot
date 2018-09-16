package config

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/config"
	"os"
	"path"
	"testing"
)

func createTestFilePath(fileName string) string {
	testDirPath, _ := os.Getwd()
	return path.Join(testDirPath, fileName)
}

func Test設定ファイルtomlを読んで値が読める(t *testing.T) {
	testFilePath := createTestFilePath("testSettings.toml")

	result := config.LoadSettings(testFilePath)

	bs := result.Bot
	assertSettingFileNumber(bs.IntervalSecond, 999, t)
	assertSettingFileNumber(bs.FrequentThresholdHour, 666, t)
	ts := result.Twitter
	assertSettingFileValue(ts.ConsumerKey, "ck", t)
	assertSettingFileValue(ts.ConsumerSecret, "cs", t)
	assertSettingFileValue(ts.AccessToken, "at", t)
	assertSettingFileValue(ts.AccessTokenSecret, "ats", t)
}

// utility functions

func assertSettingFileValue(actual string, expect string, t *testing.T) {
	if actual != expect {
		t.Errorf("設定ファイルの値が異なります。結果:%v,期待:%v", actual, expect)
	}
}

func assertSettingFileNumber(actual int, expect int, t *testing.T) {
	if actual != expect {
		t.Errorf("設定ファイルの値が異なります。結果:%d,期待:%d", actual, expect)
	}
}
