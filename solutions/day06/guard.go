package day06

import (
	"github.com/dvan-sqsp/advent-of-code-2024/util"
)

type direction int

const (
	Up direction = iota
	Right
	Down
	Left
)

type Guard struct {
	util.Position
	direction      direction
	numDistinctPos int
	visited        map[util.Position]bool
}

func NewGuard(x int, y int) *Guard {
	visited := make(map[util.Position]bool)
	visited[util.Position{X: x, Y: y}] = true
	return &Guard{
		Position:       util.Position{X: x, Y: y},
		direction:      Up, // always start up
		numDistinctPos: 1,  // starting pos is inclusive
		visited:        visited,
	}
}

func (g *Guard) HasLeftMap(patrolMap [][]string) bool {
	// determine if we hit a "#"
	nextX := g.X
	nextY := g.Y

	// preemptively check the next position
	switch g.direction {
	case Up:
		nextY--
	case Right:
		nextX++
	case Down:
		nextY++
	case Left:
		nextX--
	}

	// we're free!
	if !util.IsInBounds(patrolMap, nextX, nextY) {
		return true
	}

	// we're still in bounds, check if we hit something blocking the guard
	// if so, turn
	nextPos := patrolMap[nextY][nextX]
	if nextPos == "#" {
		g.Turn90()
	}

	// actually take the action to move
	switch g.direction {
	case Up:
		g.MoveUp()
	case Right:
		g.MoveRight()
	case Down:
		g.MoveDown()
	case Left:
		g.MoveLeft()
	}

	if _, ok := g.visited[util.Position{X: g.X, Y: g.Y}]; !ok {
		g.visited[util.Position{X: g.X, Y: g.Y}] = true
		g.numDistinctPos++
	}

	return false
}

func (g *Guard) MoveUp() {
	g.Y--
}

func (g *Guard) MoveDown() {
	g.Y++
}

func (g *Guard) MoveRight() {
	g.X++
}

func (g *Guard) MoveLeft() {
	g.X--
}

func (g *Guard) Turn90() {
	g.direction++
	if g.direction > Left {
		g.direction = Up
	}

}
