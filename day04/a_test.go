package day04

import (
	"testing"

	"github.com/Baz00k/advent-of-code-2023/helpers"
)

func TestSolvePartA(t *testing.T) {
	input, err := helpers.ReadLines("test.txt")

	if err != nil {
		panic(err)
	}

	result := SolvePartA(input)

	if result != 13 {
		t.Errorf("Expected 13, got %d", result)

		return
	}

	t.Log("Success")
}
