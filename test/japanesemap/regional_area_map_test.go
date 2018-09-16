package japanesemap

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/japanesemap"
	"testing"
)

func Test北海道沿岸の海の座標を指定しても北海道地方と判定することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalAreaMap()

	actual := sut.GetRegionName(45.07497, 144.60204)

	if actual != "北海道" {
		t.Errorf("北海道沿岸の海の座標を指定しても、'北海道'と返ってこない。値:'%v`", actual)
	}
}

func Test沖縄沿岸の海の座標を指定しても沖縄地方と判定することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalAreaMap()

	actual := sut.GetRegionName(26.6, 129.8)

	if actual != "沖縄" {
		t.Errorf("沖縄沿岸の海の座標を指定しても、'沖縄'と返ってこない。値:'%v`", actual)
	}
}

func Test仙台沿岸の海の座標を指定しても東北地方と判定することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalAreaMap()

	actual := sut.GetRegionName(37.28634, 146.15317)

	if actual != "東北" {
		t.Errorf("仙台沿岸の海の座標を指定しても、'東北'と返ってこない。値:'%v`", actual)
	}
}

func Test横浜沿岸の海の座標を指定しても関東地方と判定することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalAreaMap()

	actual := sut.GetRegionName(33.85279, 141.73036)

	if actual != "関東" {
		t.Errorf("横浜沿岸の海の座標を指定しても、'関東'と返ってこない。値:'%v`", actual)
	}
}
