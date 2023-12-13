package day05

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

	if result != 35 {
		t.Errorf("Expected 35, got %d", result)

		return
	}

	t.Log("Success")
}
