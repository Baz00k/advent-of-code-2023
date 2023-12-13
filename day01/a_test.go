package day01

import (
	"testing"

	"github.com/Baz00k/advent-of-code-2023/helpers"
)

func TestSolvePartA(t *testing.T) {
	input, err := helpers.ReadLines("test_a.txt")

	if err != nil {
		t.Fatal(err)
	}

	result := SolvePartA(input)

	if result != 142 {
		t.Errorf("Expected 142, got %d", result)

		return
	}

	t.Log("Success")
}
