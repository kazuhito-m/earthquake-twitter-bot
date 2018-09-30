package japanesemap

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/japanesemap"
	"testing"
)

func Test47都道府県の要素が詰まったオブジェクトを取得することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalClassifications()

	actual := len(sut.Classifications)
	if actual != 47 {
		t.Errorf("47都道府県分の仕様が無かった。要素数:%v", actual)
	}

	//　サンプルで一つ程度確認
	notExists := true
	for _, v := range sut.Classifications {
		if v.PrefectureName != "大阪府" {
			continue
		}
		if v.RegionName != "近畿" {
			t.Errorf("期待した地方と異なった。大阪が近畿地方ではありません。地方名:%v", v.RegionName)
		}
		notExists = false
	}
	if notExists {
		t.Errorf("都道府県に期待したものがなkった。大阪がありませんでした。")
	}
}

func Test47都道府県名から地域名が割り出すことが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalClassifications()

	assertGetRegionName("大阪府", "近畿", sut, t)
	assertGetRegionName("沖縄県", "沖縄", sut, t)
}

// Utility functions

func assertGetRegionName(prefectureName string, expect string, sut japanesemap.RegionalClassifications, t *testing.T) {
	actual := sut.GetRegionName(prefectureName)
	if actual != expect {
		t.Errorf("都道府県に対し、期待した地域名が違います。結果:%v,期待:%v", actual, expect)
	}
}
