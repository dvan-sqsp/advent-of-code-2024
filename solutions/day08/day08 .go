package day08

import (
	"fmt"
	"sync"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
	"github.com/dvan-sqsp/advent-of-code-2024/util"
)

type Day08 struct{}

var _ solver.Solver = (*Day08)(nil)

func New() *Day08 {
	return &Day08{}
}

func (d *Day08) Part1(lines []string) string {
	grid := util.Build2DMap(lines)

	antennasWithFreqs := d.buildAntennas(grid)

	uniqueAntiNodes := d.GenerateAntiNodes(grid, antennasWithFreqs, false)

	return fmt.Sprintf("%d", len(uniqueAntiNodes))
}

func (d *Day08) Part2(lines []string) string {
	grid := util.Build2DMap(lines)

	antennasWithFreqs := d.buildAntennas(grid)

	allUniqueNodes := d.GenerateAntiNodes(grid, antennasWithFreqs, true)
	return fmt.Sprintf("%d", len(allUniqueNodes))
}

func (d *Day08) buildAntennas(grid [][]string) map[string][]*Antenna {
	antennasWithFreqs := make(map[string][]*Antenna)
	for y, row := range grid {
		for x, frequency := range row {
			if frequency == "." {
				continue
			}
			if _, ok := antennasWithFreqs[frequency]; !ok {
				antennasWithFreqs[frequency] = make([]*Antenna, 0)
			}
			antennasWithFreqs[frequency] = append(antennasWithFreqs[frequency], NewAntenna(x, y, frequency))
		}
	}
	return antennasWithFreqs
}

func (d *Day08) GenerateAntiNodes(grid [][]string, antennasWithFreqs map[string][]*Antenna, isResonantHarmonics bool) map[Position]*Antenna {
	// for each frequency, calculate the potential anti-nodes possible with each other
	// ignore any that are out of bounds
	uniqueAntiNode := make(map[Position]*Antenna)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, antennas := range antennasWithFreqs {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mu *sync.Mutex, antennas []*Antenna) {
			defer wg.Done()

			for i, baseA := range antennas {
				j := 0
				for j < len(antennas) {
					// don't compute same antennas
					if i == j {
						j++
						continue
					}

					// compute delta
					dAPos := baseA.Distance(antennas[j])
					antiNodePos := NewPos(baseA.X+dAPos.X, baseA.Y+dAPos.Y)

					if isResonantHarmonics {
						// add the antenna since it counts as anti node.
						uniqueAntiNode[baseA.Position] = baseA
						// keep generating until out of bounds
						// we already know the delta so keep using it
						for util.IsInBounds(grid, antiNodePos.X, antiNodePos.Y) {
							antiNode := NewAntenna(antiNodePos.X, antiNodePos.Y, "#")
							mu.Lock()
							uniqueAntiNode[antiNodePos] = antiNode
							mu.Unlock()

							antiNodePos.X += dAPos.X
							antiNodePos.Y += dAPos.Y
						}
					} else if util.IsInBounds(grid, antiNodePos.X, antiNodePos.Y) {
						mu.Lock()
						uniqueAntiNode[antiNodePos] = NewAntenna(antiNodePos.X, antiNodePos.Y, "#")
						mu.Unlock()
					}
					j++
				}
			}
		}(&wg, &mu, antennas)
		wg.Wait()
	}
	return uniqueAntiNode
}
