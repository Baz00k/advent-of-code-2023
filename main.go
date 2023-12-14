package main

import (
	"flag"
	"log"

	"github.com/Baz00k/advent-of-code-2023/helpers"
)

func main() {
	day := flag.Int("day", 0, "Specify the day to solve.")
	d := flag.Int("d", 0, "Specify the day to solve.")
	op := flag.String("op", "solve", "Specify the operation (solve or scaffold).")
	flag.Parse()

	if *day == 0 && *d == 0 {
		log.Fatal("No day provided.")
	}

	dayToSolve := *day
	if *d != 0 {
		dayToSolve = *d
	}

	// Manually parse non-flag arguments
	args := flag.Args()
	if len(args) > 0 {
		*op = args[0]
	}

	switch *op {
	case "solve":
		Solve(dayToSolve)
	case "scaffold":
		helpers.Scaffold(dayToSolve)
	default:
		log.Fatal("Invalid operation provided.")
	}
}
