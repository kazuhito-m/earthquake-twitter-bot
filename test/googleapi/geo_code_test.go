package googleapi

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/googleapi"
	"github.com/kazuhito-m/earthquake-twitter-bot/test"
	"testing"
)

func TestGeoCode情報から主要な都道府県名を含む場所のテキストリストを取得できる(t *testing.T) {
	jsonTextForTest := test.LoadTestJson("testGoogleMapApiSampleGeoCodeRespons_hokkaiido.json")
	sut := googleapi.ParseJsonOf(jsonTextForTest)

	actual := sut.MajorPlacesIncludingPrefectureName()

	expects := []string{
		"R222+22 日本、北海道厚真町",
		"北海道",
		"Unnamed Road, 幌内 厚真町 勇払郡 北海道 059-1616 日本",
		"北海道",
		"日本 〒059-1616",
		"北海道",
		"日本、〒059-1616 北海道勇払郡厚真町幌内",
		"北海道",
		"日本、北海道勇払郡厚真町",
		"北海道",
		"日本、北海道勇払郡",
		"北海道",
		"日本、北海道",
		"日本",
	}

	actualCount := len(actual)
	expectCount := len(expects)
	if actualCount != expectCount {
		t.Errorf("主要な都道府県名を含む場所のテキストリストの要素数が期待どおりじゃない。期待:'%v',結果:'%v'", expectCount, actualCount)
	}

	for i, actual := range actual {
		expect := expects[i]
		if actual != expect {
			t.Errorf("主要な都道府県名を含む場所のテキストリストの要素が期待どおりじゃない。要素数:%v個目 ,期待:'%v',結果:'%v'", i, expect, actual)
		}
	}
}
