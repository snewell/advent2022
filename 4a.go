package main

import (
	"io"

	"github.com/snewell/advent2022/internal/aoc"
	"github.com/snewell/advent2022/internal/solutions"
)

func containedBy(first solutions.ClearingAssignment, second solutions.ClearingAssignment) bool {
	if first.Low <= second.Low && first.High >= second.High {
		return true
	}
	return false
}

func main() {
	containedCount := 0
	handler := func(first solutions.ClearingAssignment, second solutions.ClearingAssignment) {
		if containedBy(first, second) || containedBy(second, first) {
			containedCount++
		}
	}
	aoc.Run(func(r io.Reader) interface{} {
		solutions.ReadClearings(r, handler)
		return containedCount
	})
}
