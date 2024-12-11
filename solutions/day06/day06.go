package day06

import (
	"fmt"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
	"github.com/dvan-sqsp/advent-of-code-2024/util"
)

type Day06 struct {
}

var _ solver.Solver = (*Day06)(nil)

func New() *Day06 {
	return &Day06{}
}

func (d *Day06) Part1(lines []string) string {
	total := 0
	patrolMap := util.Build2DMap(lines, func(e string) string {
		return e
	})

	var guard *Guard
	for y, row := range patrolMap {
		for x, char := range row {
			if char == "^" {
				guard = NewGuard(x, y)
				break
			}
		}
	}

	for guard != nil && !guard.HasLeftMap(patrolMap) {
		total = guard.numDistinctPos
	}

	return fmt.Sprintf("%d", total)
}

func (d *Day06) Part2(lines []string) string {
	total := 0

	return fmt.Sprintf("%d", total)
}
