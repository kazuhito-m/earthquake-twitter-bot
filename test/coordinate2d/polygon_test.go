package coordinate2d

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/coordinate2d"
	"testing"
)

func Test単純な一片を1の長さとする三角形の中に点があるかを検査することが出来る(t *testing.T) {
	sut := coordinate2d.CreatePolygon([][2] float64{
		{0, 0},
		{0, 1},
		{1, 0},
	})

	assertInsideOnSquare(0.5, 0, true, sut, t)
	assertInsideOnSquare(-0.2, 0.5, false, sut, t)
	assertInsideOnSquare(0.2, 0.2, true, sut, t)
}

func Test単純な2x2の正方形の中に点があるかを検査することが出来る(t *testing.T) {
	const l = 3
	sut := coordinate2d.CreatePolygon([][2] float64{
		{0, 0},
		{0, l},
		{l, l},
		{l, 0},
	})

	assertInsideOnSquare(-1, -1, false, sut, t)
	assertInsideOnSquare(0, 0, true, sut, t)
	assertInsideOnSquare(1, 0, true, sut, t)
	assertInsideOnSquare(2, 0, true, sut, t)
	assertInsideOnSquare(3, 0, false, sut, t)
	assertInsideOnSquare(0, 1, true, sut, t)
	assertInsideOnSquare(1, 1, true, sut, t)
	assertInsideOnSquare(2, 1, true, sut, t)
	assertInsideOnSquare(3, 1, false, sut, t)
	assertInsideOnSquare(0, 2, true, sut, t)
	assertInsideOnSquare(1, 2, true, sut, t)
	assertInsideOnSquare(2, 2, true, sut, t)
	assertInsideOnSquare(3, 2, false, sut, t)
	assertInsideOnSquare(0, 3, false, sut, t)
	assertInsideOnSquare(1, 3, false, sut, t)
	assertInsideOnSquare(2, 3, false, sut, t)
	assertInsideOnSquare(3, 3, false, sut, t)
}

func Test単純な0_3までで描くの八角形の中に点が在るかを検査することが出来る(t *testing.T) {
	sut := coordinate2d.CreatePolygon([][2] float64{
		{0, 1},
		{0, 2},
		{1, 3},
		{2, 3},
		{3, 2},
		{3, 1},
		{2, 0},
		{1, 0},
	})

	assertInsideOnSquare(-1, -1, false, sut, t)
	assertInsideOnSquare(0, 0, false, sut, t)
	assertInsideOnSquare(1, 0, true, sut, t)
	assertInsideOnSquare(2, 0, false, sut, t)
	assertInsideOnSquare(3, 0, false, sut, t)
	assertInsideOnSquare(4, 0, false, sut, t)
	assertInsideOnSquare(0, 1, true, sut, t)
	assertInsideOnSquare(1, 1, true, sut, t)
	assertInsideOnSquare(2, 1, true, sut, t)
	assertInsideOnSquare(3, 1, false, sut, t)
	assertInsideOnSquare(4, 1, false, sut, t)
	assertInsideOnSquare(0, 2, true, sut, t)
	assertInsideOnSquare(1, 2, true, sut, t)
	assertInsideOnSquare(2, 2, true, sut, t)
	assertInsideOnSquare(3, 2, false, sut, t)
	assertInsideOnSquare(4, 2, false, sut, t)
	assertInsideOnSquare(0, 3, false, sut, t)
	assertInsideOnSquare(1, 3, false, sut, t)
	assertInsideOnSquare(2, 3, false, sut, t)
	assertInsideOnSquare(3, 3, false, sut, t)
	assertInsideOnSquare(4, 3, false, sut, t)
}

func Test単純な0_3をちょっと超えるくらいで描くの八角形の中に点が在るかを検査することが出来る(t *testing.T) {
	// 「未満」で判定するので「1.1など"ちょっと上"」にして「含む」を確認する。
	sut := coordinate2d.CreatePolygon([][2] float64{
		{0, 0.9},
		{0, 2.1},
		{1, 3.1},
		{2.1, 3.1},
		{3.1, 2.1},
		{3.1, 1.1},
		{2.1, 0},
		{1, 0},
	})

	assertInsideOnSquare(-1, -1, false, sut, t)
	assertInsideOnSquare(0, 0, false, sut, t)
	assertInsideOnSquare(1, 0, true, sut, t)
	assertInsideOnSquare(2, 0, true, sut, t)
	assertInsideOnSquare(3, 0, false, sut, t)
	assertInsideOnSquare(4, 0, false, sut, t)
	assertInsideOnSquare(0, 1, true, sut, t)
	assertInsideOnSquare(1, 1, true, sut, t)
	assertInsideOnSquare(2, 1, true, sut, t)
	assertInsideOnSquare(3, 1, true, sut, t)
	assertInsideOnSquare(4, 1, false, sut, t)
	assertInsideOnSquare(0, 2, true, sut, t)
	assertInsideOnSquare(1, 2, true, sut, t)
	assertInsideOnSquare(2, 2, true, sut, t)
	assertInsideOnSquare(3, 2, true, sut, t)
	assertInsideOnSquare(4, 2, false, sut, t)
	assertInsideOnSquare(0, 3, false, sut, t)
	assertInsideOnSquare(1, 3, true, sut, t)
	assertInsideOnSquare(2, 3, true, sut, t)
	assertInsideOnSquare(3, 3, false, sut, t)
	assertInsideOnSquare(4, 3, false, sut, t)
}

func Test論理演算でXORを検査することが出来る(t *testing.T) {
	xorPattern := [4][3]bool{
		{true, true, false},
		{true, false, true},
		{false, true, true},
		{false, false, false},
	}
	for _, v := range xorPattern {
		actual := coordinate2d.Xor(v[0], v[1])
		if (actual != v[2]) {
			t.Errorf("XORが期待通りの値に鳴らない。%v ^ %v = %v", v[0], v[1], v[2])
		}
	}
}

// Utility functions

func assertInsideOnSquare(X float64, Y float64, expect bool, sut coordinate2d.Polygon, t *testing.T) {
	target := coordinate2d.Point{X, Y}
	actual := sut.InsideOf(target)
	result := "入ってない"
	if actual {
		result = "入ってる"
	}
	//fmt.Printf("(%v,%v)%v\n", target.X, target.Y, result)
	if actual != expect {
		t.Errorf("長さ2の辺を持つ正方形の中に座標(%v,%v)の点は「%v」判定となった。", target.X, target.Y, result)
	}
}
