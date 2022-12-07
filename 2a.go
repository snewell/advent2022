package main

import (
	"io"

	"github.com/snewell/advent2022/internal/aoc"
	"github.com/snewell/advent2022/internal/solutions"
)

func main() {
	score := 0
	handler := func(opponent solutions.RpsSign, you solutions.RpsSign) {
		result := solutions.CheckRpsRound(opponent, you)
		score += int(result) + int(you)
	}
	aoc.Run(func(r io.Reader) interface{} {
		solutions.PlayRps(r, handler)
		return score
	})
}
