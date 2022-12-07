package main

import (
	"io"

	"github.com/snewell/advent2022/internal/aoc"
	"github.com/snewell/advent2022/internal/solutions"
)

func findIndex(data [3]int, newCalories int) int {
	for i := 2; i >= 0; i-- {
		if newCalories > data[i] {
			return i
		}
	}
	return -1
}

func main() {
	maxCalories := [3]int{0, 0, 0}
	handler := func(calories int) {
		index := findIndex(maxCalories, calories)
		if index != -1 {
			for i := 0; i < index; i++ {
				maxCalories[i] = maxCalories[i+1]
			}
			maxCalories[index] = calories
		}
	}
	aoc.Run(func(r io.Reader) interface{} {
		solutions.CountCalories(r, handler)
		return maxCalories[0] + maxCalories[1] + maxCalories[2]
	})
}
