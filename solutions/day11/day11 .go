package day11

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
)

const (
	part1Blinks = 25
	part2Blinks = 75
)

type Day11 struct{}

var _ solver.Solver = (*Day11)(nil)

func New() *Day11 {
	return &Day11{}
}

func (d *Day11) Part1(lines []string) string {
	// Just use a linked list because I think it is neat
	stones := list.New()
	for _, stoneStr := range strings.Split(lines[0], " ") {
		stoneNum, _ := strconv.Atoi(stoneStr)
		stones.PushBack(NewStone(stoneNum))
	}

	// blinkLinkedList
	for range part1Blinks {
		d.blinkLinkedList(stones)
	}
	return fmt.Sprintf("%d", stones.Len())
}

func (d *Day11) Part2(lines []string) string {
	// keep this map updated with new stones
	stones := make(map[Stone]int)

	// so, order doesn't really matter
	// we can just compute each stone and keep track of the count
	// of how many times the stone appears
	for _, stoneStr := range strings.Split(lines[0], " ") {
		stoneNum, _ := strconv.Atoi(stoneStr)
		stone := NewStone(stoneNum)
		stones[*stone] = 1
	}

	// blinkLinkedList
	for range part2Blinks {
		stones = d.blinkHashMap(stones)
	}

	total := 0
	for _, count := range stones {
		total += count
	}
	return fmt.Sprintf("%d", total)
}

func (d *Day11) blinkLinkedList(stones *list.List) {
	for e := stones.Front(); e != nil; e = e.Next() {
		stone := e.Value.(*Stone)
		// three rules
		// 1. if 0, change to 1
		// 2. if even num string length, split
		// 3. else multiply by 2024
		if stone.value == 0 {
			stone.RuleOne()
		} else if stone.IsNumberLengthEven() {
			// 1 -> 2425 -> 3
			// 24 25
			newStone1, newStone2 := stone.RuleTwo()

			// 1 -> 24 -> 2425 -> 3
			// 1 -> 24 -> 25 -> 2425 -> 3
			// 1 -> 24 -> 25 -> 3
			stones.InsertBefore(newStone1, e)
			nextE := stones.InsertAfter(newStone2, e)
			stones.Remove(e)

			// update the current element to the last stone we inserted
			e = nextE
		} else {
			stone.RuleThree()
		}
	}
}

func (d *Day11) blinkHashMap(stones map[Stone]int) map[Stone]int {
	// keep track of new stones we create in this blink iteration
	newStones := make(map[Stone]int)

	// keep track of new stones generated
	incCount := func(key Stone, incr int) {
		if _, ok := newStones[key]; !ok {
			newStones[key] = 0
		}
		newStones[key] += incr
	}

	for stone, count := range stones {
		// three rules
		// 1. if 0, change to 1
		// 2. if even num string length, split
		// 3. else multiply by 2024
		if stone.value == 0 {
			incCount(*stone.RuleOne(), count)
		} else if stone.IsNumberLengthEven() {
			s1, s2 := stone.RuleTwo()
			incCount(*s1, count)
			incCount(*s2, count)
		} else {
			incCount(*stone.RuleThree(), count)
		}
	}

	return newStones
}
