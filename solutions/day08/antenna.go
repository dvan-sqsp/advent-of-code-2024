package day08

import (
	"github.com/dvan-sqsp/advent-of-code-2024/util"
)

type Antenna struct {
	util.Position
	Frequency string
}

func NewAntenna(x int, y int, frequency string) *Antenna {
	return &Antenna{
		Position: NewPos(x, y), Frequency: frequency,
	}
}

func NewPos(x, y int) util.Position {
	return util.Position{X: x, Y: y}
}

func (a *Antenna) Distance(b *Antenna) util.Position {
	return util.Position{
		X: a.X - b.Position.X,
		Y: a.Y - b.Position.Y,
	}
}
