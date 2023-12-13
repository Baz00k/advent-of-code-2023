package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Baz00k/advent-of-code-2023/day01"
	"github.com/Baz00k/advent-of-code-2023/day02"
	"github.com/Baz00k/advent-of-code-2023/day03"
	"github.com/Baz00k/advent-of-code-2023/day04"
	"github.com/Baz00k/advent-of-code-2023/day05"
)

func main() {
	day := flag.Int("day", 0, "Specify the day to solve.")
	d := flag.Int("d", 0, "Specify the day to solve.")
	flag.Parse()

	if *day == 0 && *d == 0 {
		log.Fatal("No day provided.")
	}

	dayToSolve := *day
	if *d != 0 {
		dayToSolve = *d
	}

	switch dayToSolve {
	case 1:
		day01.Solve()
	case 2:
		day02.Solve()
	case 3:
		day03.Solve()
	case 4:
		day04.Solve()
	case 5:
		day05.Solve()
	default:
		fmt.Println("Invalid day provided.")
	}
}
