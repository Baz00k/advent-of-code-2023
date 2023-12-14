package day06

import (
	"testing"

	"github.com/Baz00k/advent-of-code-2023/helpers"
)

func TestSolvePartA(t *testing.T) {
	input, err := helpers.ReadLines("test.txt")

	if err != nil {
		t.Fatal(err)
	}

	result := SolvePartA(input)

	if result != 288 {
		t.Errorf("Expected 288, got %d", result)

		return
	}

	t.Log("Success")
}
