package day06

import (
	"fmt"

	"github.com/Baz00k/advent-of-code-2023/helpers"
)

func Solve() {
	input, err := helpers.ReadLines("day06/input.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("Day 06")
	fmt.Println("Part A:", SolvePartA(input))
	fmt.Println("Part B:", SolvePartB(input))
}
