package day03

import (
	"testing"

	"github.com/Baz00k/advent-of-code-2023/helpers"
)

func TestSolvePartB(t *testing.T) {
	input, err := helpers.ReadChars("test.txt")

	if err != nil {
		panic(err)
	}

	result := SolvePartB(input)

	if result != 467835 {
		t.Errorf("Expected 467835, got %d", result)

		return
	}

	t.Log("Success")
}