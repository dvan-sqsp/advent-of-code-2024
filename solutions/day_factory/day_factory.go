package day_factory

import (
	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
	"github.com/dvan-sqsp/advent-of-code-2024/solutions/day01"
	"github.com/dvan-sqsp/advent-of-code-2024/solutions/day02"
	"github.com/dvan-sqsp/advent-of-code-2024/solutions/day03"
	"github.com/dvan-sqsp/advent-of-code-2024/solutions/day04"
	"github.com/dvan-sqsp/advent-of-code-2024/solutions/day05"
	"github.com/dvan-sqsp/advent-of-code-2024/solutions/day06"
	"github.com/dvan-sqsp/advent-of-code-2024/solutions/day07"
	"github.com/dvan-sqsp/advent-of-code-2024/solutions/day08"
	"github.com/dvan-sqsp/advent-of-code-2024/solutions/day09"
	"github.com/dvan-sqsp/advent-of-code-2024/solutions/day10"
	"github.com/rs/zerolog/log"
)

func GetDay(day int) solver.Solver {
	// for each day, add a new solver
	switch day {
	case 1:
		return day01.New()
	case 2:
		return day02.New()
	case 3:
		return day03.New()
	case 4:
		return day04.New()
	case 5:
		return day05.New()
	case 6:
		return day06.New()
	case 7:
		return day07.New()
	case 8:
		return day08.New()
	case 9:
		return day09.New()
	case 10:
		return day10.New()
	default:
		log.Error().Msg("Invalid Day Number")
		return nil
	}
}
