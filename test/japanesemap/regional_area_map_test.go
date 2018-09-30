package japanesemap

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/japanesemap"
	"testing"
)

func Test北海道沿岸の海の座標を指定しても北海道地方と判定することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalAreaMap()

	actual := sut.GetRegionName(45.07497, 144.60204)

	if actual != "北海道" {
		t.Errorf("北海道沿岸の海の座標を指定しても、'北海道'と返ってこない。値:'%v'", actual)
	}
}

func Test沖縄沿岸の海の座標を指定しても沖縄地方と判定することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalAreaMap()

	actual := sut.GetRegionName(26.6, 129.8)

	if actual != "沖縄" {
		t.Errorf("沖縄沿岸の海の座標を指定しても、'沖縄'と返ってこない。値:'%v'", actual)
	}
}

func Test仙台沿岸の海の座標を指定しても東北地方と判定することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalAreaMap()

	actual := sut.GetRegionName(37.28634, 146.15317)

	if actual != "東北" {
		t.Errorf("仙台沿岸の海の座標を指定しても、'東北'と返ってこない。値:'%v'", actual)
	}
}

func Test横浜沿岸の海の座標を指定しても関東地方と判定することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalAreaMap()

	actual := sut.GetRegionName(33.85279, 141.73036)

	if actual != "関東" {
		t.Errorf("横浜沿岸の海の座標を指定しても、'関東'と返ってこない。値:'%v'", actual)
	}
}

func Test石川沿岸の海の座標を指定しても中部地方と判定することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalAreaMap()

	actual := sut.GetRegionName(36.58646, 136.00005)

	if actual != "中部" {
		t.Errorf("石川沿岸の海の座標を指定しても、'中部'と返ってこない。値:'%v'", actual)
	}
}

func Test和歌山沿岸の海の座標を指定しても関西地方と判定することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalAreaMap()

	actual := sut.GetRegionName(31.8334, 137.68597)

	if actual != "関西" {
		t.Errorf("和歌山沿岸の海の座標を指定しても、'関西'と返ってこない。値:'%v'", actual)
	}
}

func Test四国沿岸の海の座標を指定しても四国地方と判定することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalAreaMap()

	actual := sut.GetRegionName(30.28753, 135.54883)

	if actual != "四国" {
		t.Errorf("四国沿岸の海の座標を指定しても、'四国'と返ってこない。値:'%v'", actual)
	}
}

func Test島根沿岸の海の座標を指定しても中国地方と判定することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalAreaMap()

	actual := sut.GetRegionName(35.54408, 131.13412)

	if actual != "中国" {
		t.Errorf("島根沿岸の海の座標を指定しても、'中国'と返ってこない。値:'%v'", actual)
	}
}


func Test鹿児島沿岸の海の座標を指定しても九州地方と判定することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalAreaMap()

	actual := sut.GetRegionName(28.81353, 132.67783)

	if actual != "九州" {
		t.Errorf("鹿児島沿岸の海の座標を指定しても、'九州'と返ってこない。値:'%v'", actual)
	}
}
