package day03

import (
	"testing"

	"github.com/Baz00k/advent-of-code-2023/helpers"
)

func TestSolvePartA(t *testing.T) {
	input, err := helpers.ReadChars("test.txt")

	if err != nil {
		panic(err)
	}

	result := SolvePartA(input)

	if result != 4361{
		t.Errorf("Expected 4361, got %d", result)

		return
	}

	t.Log("Success")
}