package main

import (
	"io"

	"github.com/snewell/advent2022/internal/aoc"
	"github.com/snewell/advent2022/internal/solutions"
)

func main() {
	score := 0
	resultMap := [3][2]solutions.RpsSign{
		{solutions.ScissorsSign, solutions.PaperSign},
		{solutions.RockSign, solutions.ScissorsSign},
		{solutions.PaperSign, solutions.RockSign},
	}
	handler := func(opponent solutions.RpsSign, result solutions.RpsResult) {
		var you solutions.RpsSign
		switch result {
		case solutions.RpsDraw:
			you = opponent

		case solutions.RpsLoss:
			you = resultMap[int(opponent)-1][0]

		case solutions.RpsWin:
			you = resultMap[int(opponent)-1][1]
		}
		score += int(result) + int(you)
	}
	aoc.Run(func(r io.Reader) interface{} {
		solutions.ExpectRps(r, handler)
		return score
	})
}
