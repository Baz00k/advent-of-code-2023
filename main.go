package main

import (
	"flag"
	"fmt"

	"github.com/Baz00k/advent-of-code-2023/day01"
	"github.com/Baz00k/advent-of-code-2023/day02"
	"github.com/Baz00k/advent-of-code-2023/day03"
)

func main() {
    day := flag.Int("day", 0, "Specify the day to solve.")
    d := flag.Int("d", 0, "Specify the day to solve.")
    flag.Parse()

    if *day == 0 && *d == 0 {
        fmt.Println("Please provide a day to solve.")
        return
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
    default:
        fmt.Println("Invalid day provided.")
    }
}