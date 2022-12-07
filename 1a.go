package main

import (
	"io"

	"github.com/snewell/advent2022/internal/aoc"
	"github.com/snewell/advent2022/internal/solutions"
)

func main() {
	maxCalories := 0
	handler := func(calories int) {
		if calories > maxCalories {
			maxCalories = calories
		}
	}
	aoc.Run(func(r io.Reader) interface{} {
		solutions.CountCalories(r, handler)
		return maxCalories
	})
}
