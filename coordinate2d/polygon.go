package coordinate2d

type Polygon struct {
	Points []Point
}

func (p Polygon) InsideOf(target Point) bool {
	points := p.Points
	inside := false
	var cp Point
	for i, nowPoint := range points {
		previousPoint := previousPoint(i, points)
		if !cross(target.Y, nowPoint, previousPoint, &cp) {
			continue
		}
		if cp.X > target.X {
			inside = !inside
		}
	}
	return inside;
}

func cross(y float64, p1 Point, p2 Point, cp *Point) bool {
	if p1.Y > p2.Y {
		p := p1
		p1 = p2
		p2 = p
	}
	cp.X = (p1.X*(p2.Y-y) + p2.X*(y-p1.Y)) / (p2.Y - p1.Y);
	cp.Y = y;
	return p1.Y <= y && y < p2.Y
}

func previousPoint(index int, points []Point) Point {
	previousIndex := (index + 1) % len(points)
	return points[previousIndex]
}

func CreatePolygon(pointArray [][2] float64) Polygon {
	count := len(pointArray)
	points := make([]Point, count)
	for i, v := range pointArray {
		points[i] = Point{v[0], v[1]}
	}
	return Polygon{points}
}
