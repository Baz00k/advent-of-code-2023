package main

import (
	"log"

	"github.com/Baz00k/advent-of-code-2023/day01"
	"github.com/Baz00k/advent-of-code-2023/day02"
	"github.com/Baz00k/advent-of-code-2023/day03"
	"github.com/Baz00k/advent-of-code-2023/day04"
	"github.com/Baz00k/advent-of-code-2023/day05"
	"github.com/Baz00k/advent-of-code-2023/day06"
	"github.com/Baz00k/advent-of-code-2023/day07"
)

type SolveFunc func()

var solveFuncs = map[int]SolveFunc{
	1: day01.Solve,
	2: day02.Solve,
	3: day03.Solve,
	4: day04.Solve,
	5: day05.Solve,
	6: day06.Solve,
	7: day07.Solve,
}

func Solve(day int) {
	if solve, ok := solveFuncs[day]; ok {
		solve()
	} else {
		log.Fatalf("Day %d not implemented.", day)
	}
}
