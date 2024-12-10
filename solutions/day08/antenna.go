package day08

type Antenna struct {
	Position
	Frequency string
}

type Position struct {
	X, Y int
}

func NewAntenna(x int, y int, frequency string) *Antenna {
	return &Antenna{
		Position: NewPos(x, y), Frequency: frequency,
	}
}

func NewPos(x, y int) Position {
	return Position{X: x, Y: y}
}

func (a *Antenna) Distance(b *Antenna) Position {
	return Position{
		X: a.X - b.Position.X,
		Y: a.Y - b.Position.Y,
	}
}
