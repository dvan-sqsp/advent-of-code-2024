package day10

import (
	"fmt"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
	"github.com/dvan-sqsp/advent-of-code-2024/util"
)

type Day10 struct{}

var _ solver.Solver = (*Day10)(nil)

func New() *Day10 {
	return &Day10{}
}

type TrailHead struct {
	util.Position
	value  int
	score  int
	rating int
}

func NewTrailHead(pos util.Position, value int) *TrailHead {
	return &TrailHead{pos, value, 0, 0}
}

func (d *Day10) Part1(lines []string) string {
	grid, trailHeads := d.buildTrail(lines)

	score, _ := d.calculateScore(grid, trailHeads)

	return fmt.Sprintf("%d", score)
}

func (d *Day10) Part2(lines []string) string {
	grid, trailHeads := d.buildTrail(lines)

	_, rating := d.calculateScore(grid, trailHeads)

	return fmt.Sprintf("%d", rating)
}

func (d *Day10) buildTrail(lines []string) ([][]int, []TrailHead) {
	grid := util.Build2DMapInt(lines)

	trailHeads := make([]TrailHead, 0)
	for y, row := range grid {
		for x, val := range row {
			if val != 0 {
				continue
			}

			trailHeads = append(trailHeads, *NewTrailHead(util.Position{
				X: x,
				Y: y,
			}, val))
		}
	}

	return grid, trailHeads
}

func (d *Day10) calculateScore(grid [][]int, trailHeads []TrailHead) (int, int) {
	totalScore := 0
	totalRating := 0
	// for each trail head, DFS through until we reach all 9s possible
	for _, head := range trailHeads {
		visited := make(map[util.Position]bool)
		d.traverse(grid, &head, head.value, head.X, head.Y, visited)

		totalScore += head.score
		totalRating += head.rating
	}

	return totalScore, totalRating
}

func (d *Day10) traverse(grid [][]int, trailHead *TrailHead, curVal int, curX int, curY int, visited map[util.Position]bool) {
	pos := util.Position{X: curX, Y: curY}

	// base case, we've found a 9!
	// keep track of the position forever since we only care about the first visit
	if curVal == 9 {
		if _, found := visited[pos]; !found {
			visited[util.Position{X: curX, Y: curY}] = true
			trailHead.score++
		}
		trailHead.rating++
		return
	}

	// another case! we've visited this position already
	if visited[pos] {
		return
	}

	visited[pos] = true

	// attempt to move up
	if util.IsInBounds(grid, curX, curY-1) {
		nextVal := grid[curY-1][curX]
		if curVal+1 == nextVal {
			d.traverse(grid, trailHead, nextVal, curX, curY-1, visited)
		}
	}

	// move down
	if util.IsInBounds(grid, curX, curY+1) {
		nextVal := grid[curY+1][curX]
		if curVal+1 == nextVal {
			d.traverse(grid, trailHead, nextVal, curX, curY+1, visited)
		}
	}

	// move right
	if util.IsInBounds(grid, curX+1, curY) {
		nextVal := grid[curY][curX+1]
		if curVal+1 == nextVal {
			d.traverse(grid, trailHead, nextVal, curX+1, curY, visited)
		}
	}

	// move left
	if util.IsInBounds(grid, curX-1, curY) {
		nextVal := grid[curY][curX-1]
		if curVal+1 == nextVal {
			d.traverse(grid, trailHead, nextVal, curX-1, curY, visited)
		}
	}

	// unravelling, we can revisit now
	visited[pos] = false
}
