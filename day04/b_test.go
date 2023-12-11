package day04

import (
	"testing"

	"github.com/Baz00k/advent-of-code-2023/helpers"
)

func TestSolvePartB(t *testing.T) {
	input, err := helpers.ReadLines("test.txt")

	if err != nil {
		panic(err)
	}

	result := SolvePartB(input)

	if result != 30 {
		t.Errorf("Expected 30, got %d", result)

		return
	}

	t.Log("Success")
}
