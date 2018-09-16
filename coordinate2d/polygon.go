package coordinate2d

type Polygon struct {
	Points []Point
}

func (p Polygon) InsideOf(target Point) bool {
	points := p.Points
	inside := false
	for i, nowPoint := range points {
		previousPoint := previousPoint(i, points)
		line := Line{nowPoint, previousPoint}

		p1 := line.First()
		p2 := line.Last()
		y := target.Y
		x := (p1.X*(p2.Y-y) + p2.X*(y-p1.Y)) / (p2.Y - p1.Y);

		cross := p1.Y <= y && y < p2.Y && x > target.X

		if cross {
			inside = !inside
		}
	}
	return inside;
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
