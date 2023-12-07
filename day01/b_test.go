package day01

import (
	"testing"

	"github.com/Baz00k/advent-of-code-2023/helpers"
)

func TestSolvePartB(t *testing.T) {
	input, err := helpers.ReadLines("test_b.txt")

	if err != nil {
		panic(err)
	}

	result := SolvePartB(input)

	if result != 281 {
		t.Errorf("Expected 281, got %d", result)

		return
	}

	t.Log("Success")
}