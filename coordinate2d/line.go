package coordinate2d

type Line struct {
	p1 Point
	p2 Point
}

func (l Line) First() Point {
	if l.p1.Y > l.p2.Y {
		return l.p2
	} else {
		return l.p1
	}
}

func (l Line) Last() Point {
	if l.p1.Y > l.p2.Y {
		return l.p1
	} else {
		return l.p2
	}
}
