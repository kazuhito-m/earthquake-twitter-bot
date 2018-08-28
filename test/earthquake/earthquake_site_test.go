package earthquake

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/earthquake"
	"testing"
	"time"
)

func Test地震サイトの_時間ID_っぽいものを作る(t *testing.T) {
	actual := earthquake.CreateTimeId()
	if len(actual) != 14 {
		t.Errorf("時刻IDの値生成で文字数が想定どおりでない。値:'%v'", actual)
	}
	if !assertParseByFormat(actual, "20060102150405") {
		t.Errorf("時刻IDは変換できたが、値が日付時刻として成り立たない値だった。。値:'%v'", actual)
	}
}

func Test地震データをサイトから取得しオブジェクト化できる(t *testing.T) {
	returnJson := "{\"result\": {\"status\": \"大成功\"}}"
	client := CreateMockClient(returnJson)
	sut := earthquake.CreateSite(client)

	result := sut.GetNowEarthquakeAnalyze()

	actual := result.Result.Status
	expect := "大成功"
	if actual != expect {
		t.Errorf("取得できた地震データが期待の値ではない。期待:%v,結果:%v", actual, expect)
	}

	expectUrl := "http://www.kmoni.bosai.go.jp/new/webservice/hypo/eew/"
	if client.NoContainsExecutedUrl(expectUrl) {
		t.Errorf("実行時に期待したURLにリクエストを送っていない。期待部分(前方一致):%v,実際:%v", expectUrl, client.UrlWhenExecute())
	}
}

func Test期待以外の文字列が来た場合はステータス部分がから文字になる(t *testing.T) {
	returnJson := "JSONとして存在不能の文字列"
	client := CreateMockClient(returnJson)
	sut := earthquake.CreateSite(client)

	result := sut.GetNowEarthquakeAnalyze()

	actual := result.Result.Status
	if actual != "" {
		t.Errorf("JSONとして成り立たない文字列でるのにStatusが空文字じゃなかった。結果:%v", actual)
	}
}

func assertParseByFormat(dateTimeText string, format string) bool {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	_, error := time.ParseInLocation(format, dateTimeText, loc)
	return error == nil
}
