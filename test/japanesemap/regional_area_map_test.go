package japanesemap

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/japanesemap"
	"testing"
)

func Test北海道沿岸の海の座標を指定しても北海道地方と判定することが出来る(t *testing.T) {
	sut := japanesemap.CreateRegionalAreaMap()

	actual := sut.GetRegionName(45.07497, 144.60204)

	if actual != "北海道" {
		t.Errorf("北海道沿岸の海の座標を指定しても、'北海道'と返ってこない。値:'%v`" , actual)
	}
}
