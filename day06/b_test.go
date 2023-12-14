package day06

import (
	"testing"

	"github.com/Baz00k/advent-of-code-2023/helpers"
)

func TestSolvePartB(t *testing.T) {
	input, err := helpers.ReadLines("test.txt")

	if err != nil {
		t.Fatal(err)
	}

	result := SolvePartB(input)

	if result != 71503 {
		t.Errorf("Expected 71503, got %d", result)

		return
	}

	t.Log("Success")
}
