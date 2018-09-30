package coordinate2d

type Polygon struct {
	Points []Point
}

/*
 指定された座標が、自身が保持する「多角形」の「内側にあるか」を判定し、真偽値で帰す。
 参考: http://www.not-enough.org/abe/manual/argo/poly-naibu.html, http://home.a00.itscom.net/hatada/c01/algorithm/polygon01.html
 上記サイトのものをそのまま流用したので、意味は分かっていない。
 */
func (p Polygon) InsideOf(target Point) bool {
	points := p.Points
	inside := false
	for i, nowPoint := range points {
		previousPoint := previousPoint(i, points)
		line := Line{nowPoint, previousPoint}
		inside = Xor(inside, isCross(line, target))
	}
	return inside;
}

func isCross(line Line, target Point) bool {
	p1 := line.First()
	p2 := line.Last()
	y := target.Y
	x := (p1.X*(p2.Y-y) + p2.X*(y-p1.Y)) / (p2.Y - p1.Y)
	return p1.Y <= y && y < p2.Y && x > target.X
}

func previousPoint(index int, points []Point) Point {
	previousIndex := (index + 1) % len(points)
	return points[previousIndex]
}

func Xor(p bool, q bool) bool {
	return p != q
}

func CreatePolygon(pointArray [][2] float64) Polygon {
	count := len(pointArray)
	points := make([]Point, count)
	for i, v := range pointArray {
		points[i] = Point{v[0], v[1]}
	}
	return Polygon{points}
}
