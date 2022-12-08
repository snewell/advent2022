package main

import (
	"io"

	"github.com/snewell/advent2022/internal/aoc"
	"github.com/snewell/advent2022/internal/solutions"
)

func addToMap(items string, data *map[rune]struct{}) {
	for _, item := range items {
		(*data)[item] = struct{}{}
	}
}

func findBadge(items *[3]map[rune]struct{}) rune {
	counts := map[rune]int{}
	for index := range *items {
		for rucksackItem := range (*items)[index] {
			counts[rucksackItem]++
			if counts[rucksackItem] == 3 {
				return rucksackItem
			}
		}
	}
	panic("No common item")
	return 'a'
}

func main() {
	value := 0
	items := [3]map[rune]struct{}{
		map[rune]struct{}{},
		map[rune]struct{}{},
		map[rune]struct{}{},
	}
	count := 0
	handler := func(first string, second string) {
		addToMap(first, &items[count])
		addToMap(second, &items[count])
		count++
		if count == 3 {
			badge := findBadge(&items)
			value += solutions.GetPriority(badge)
			count = 0
			for index := range items {
				items[index] = map[rune]struct{}{}
			}
		}
	}
	aoc.Run(func(r io.Reader) interface{} {
		solutions.Rucksack(r, handler)
		return value
	})
}
