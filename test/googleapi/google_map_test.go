package googleapi

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/config"
	"github.com/kazuhito-m/earthquake-twitter-bot/googleapi"
	"github.com/kazuhito-m/earthquake-twitter-bot/test"
	"github.com/kazuhito-m/earthquake-twitter-bot/test/earthquake"
	"strings"
	"testing"
)

func TestパラメータからAPIのURLを組み立てることが出来る(t *testing.T) {
	settings := config.GoogleApiSettings{"AbCxYz"}

	actual := googleapi.CreateUrl(123.456, 987.654, settings)

	if strings.Index(actual, "https://maps.googleapis.com") != 0 {
		t.Errorf("URLの先頭が「GoogleMapAPI」のものではない。値:'%v'", actual)
	}
	if !strings.Contains(actual, "&key=AbCxYz&latlng=123.456,987.654") {
		t.Errorf("URLにキーと緯度経度が含まれていない。値:'%v'", actual)
	}
}

func TestGoogleMapAPIからJSONのデータを取得して構造体として返すことが出来る_例として日本標準時の明石天文台(t *testing.T) {
	returnJson := test.LoadTestJson("testGoogleMapApiSampleGeoCodeRespons_akashi.json")
	client := earthquake.CreateMockClient(returnJson)
	settings := config.GoogleApiSettings{"ダミーのAPIキー"}
	sut := googleapi.CreateGoogleMap(client, settings)

	result := sut.GeoCode(34.64, 135)

	if !result.Ok() {
		t.Errorf("取得出来たGoogleMapのGeoCodeのAPI成功判定が「失敗」となった。")
	}

	actual := result.Plus_Code.Compound_Code
	expect := "兵庫県明石市"
	if !strings.Contains(actual, expect) {
		t.Errorf("取得できたMapデータのcompound_codeに期待した地名が含まれていない。期待(含まれて欲しい単語):'%v',結果:'%v'", expect, actual)
	}

	count := len(result.Results[0].Address_Components)
	if count != 8 {
		t.Errorf("取得できたMapデータのfaddress_componentsの要素数が期待した個数じゃない。期待:'%v',結果:'%v'", 8, count)
	}

	actual = result.Results[0].Formatted_Address
	expect = "明石市 兵庫県"
	if !strings.Contains(actual, expect) {
		t.Errorf("取得できたMapデータのformatted_addressに期待した地名が含まれていない。期待(含まれて欲しい単語):'%v',結果:'%v'", expect, actual)
	}
}

func TestGoogleMapAPIからJSONのデータを取得して構造体として返すことが出来る_例として胆振地方中東部(t *testing.T) {
	returnJson := test.LoadTestJson("testGoogleMapApiSampleGeoCodeRespons_hokkaiido.json")
	client := earthquake.CreateMockClient(returnJson)
	settings := config.GoogleApiSettings{"ダミーのAPIキー"}
	sut := googleapi.CreateGoogleMap(client, settings)

	result := sut.GeoCode(34.64, 135)

	if !result.Ok() {
		t.Errorf("取得出来たGoogleMapのGeoCodeのAPI成功判定が「失敗」となった。")
	}

	actual := result.Plus_Code.Compound_Code
	expect := "北海道厚真町"
	if !strings.Contains(actual, expect) {
		t.Errorf("取得できたMapデータのcompound_codeに期待した地名が含まれていない。期待(含まれて欲しい単語):'%v',結果:'%v'", expect, actual)
	}

	count := len(result.Results)
	if count != 7 {
		t.Errorf("取得できたMapデータのresultsの要素数が期待した個数じゃない。期待:'%v',結果:'%v'", 7, count)
	}

	count = len(result.Results[0].Address_Components)
	if count != 7 {
		t.Errorf("取得できたMapデータのfaddress_componentsの要素数が期待した個数じゃない。期待:'%v',結果:'%v'", 7, count)
	}

	actual = result.Results[0].Formatted_Address
	expect = "厚真町 勇払郡 北海道"
	if !strings.Contains(actual, expect) {
		t.Errorf("取得できたMapデータのformatted_addressに期待した地名が含まれていない。期待(含まれて欲しい単語):'%v',結果:'%v'", expect, actual)
	}
}
