package day07

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

	if result != 0 {
		t.Errorf("Expected 0, got %d", result)

		return
	}

	t.Log("Success")
}
