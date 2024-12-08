package utils

type Pt struct {
	X, Y int
}

func (p Pt) Move(dx, dy int) Pt {
	return Pt{p.X + dx, p.Y + dy}
}

func (p Pt) Dist(p2 Pt) (dx, dy int) {
	dx = p.X - p2.X
	dy = p.Y - p2.Y
	return
}