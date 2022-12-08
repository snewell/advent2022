package main

import (
	"io"

	"github.com/snewell/advent2022/internal/aoc"
	"github.com/snewell/advent2022/internal/solutions"
)

func main() {
	value := 0
	handler := func(first string, second string) {
		for _, f := range first {
			for _, s := range second {
				if f == s {
					value += solutions.GetPriority(f)
					return
				}
			}
		}
	}
	aoc.Run(func(r io.Reader) interface{} {
		solutions.Rucksack(r, handler)
		return value
	})
}
