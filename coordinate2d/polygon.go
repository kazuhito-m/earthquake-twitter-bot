package coordinate2d

type Polygon struct {
	Points []Point
}

func (p Polygon) InsideOf(target Point) bool {
	poly := p.Points
	var j int
	count := 0
	var cp Point
	for i, _ := range poly {
		j = (i + 1) % len(poly);
		if cross(target.Y, poly[i], poly[j], &cp) {
			if cp.X > target.X {
				count++
			}
		}
	}
	return count%2 != 0;
}

func cross(y float64, p1 Point, p2 Point, cp *Point) bool {
	if p1.Y > p2.Y {
		p := p1
		p1 = p2
		p2 = p
	}
	cp.X = (p1.X*(p2.Y-y) + p2.X*(y-p1.Y)) / (p2.Y - p1.Y);
	cp.Y = y;
	return (p1.Y <= y && y < p2.Y)
}

func CreatePolygon(pointArray [][2] float64) Polygon {
	count := len(pointArray)
	points := make([]Point, count)
	for i, v := range pointArray {
		points[i] = Point{v[0], v[1]}
	}
	return Polygon{points}
}
