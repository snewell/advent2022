package main

import (
	"io"

	"github.com/snewell/advent2022/internal/aoc"
	"github.com/snewell/advent2022/internal/solutions"
)

func overlaps(first solutions.ClearingAssignment, second solutions.ClearingAssignment) bool {
	if first.Low <= second.High && first.High >= second.Low {
		return true
	}
	return false
}

func main() {
	overlapCount := 0
	handler := func(first solutions.ClearingAssignment, second solutions.ClearingAssignment) {
		if overlaps(first, second) || overlaps(second, first) {
			overlapCount++
		}
	}
	aoc.Run(func(r io.Reader) interface{} {
		solutions.ReadClearings(r, handler)
		return overlapCount
	})
}
