package day02

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

	if result != 8 {
		t.Errorf("Expected 8, got %d", result)

		return
	}

	t.Log("Success")
}