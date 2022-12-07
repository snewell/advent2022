package main

import (
	"io"

	"github.com/snewell/advent2022/internal/aoc"
	"github.com/snewell/advent2022/internal/solutions"
)

func main() {
	resultMap := [3][3]solutions.RpsResult{
		{solutions.RpsDraw, solutions.RpsWin, solutions.RpsLoss},
		{solutions.RpsLoss, solutions.RpsDraw, solutions.RpsWin},
		{solutions.RpsWin, solutions.RpsLoss, solutions.RpsDraw},
	}
	score := 0
	handler := func(opponent solutions.RpsSign, you solutions.RpsSign) {
		result := resultMap[int(opponent)-1][int(you)-1]
		score += int(result) + int(you)
	}
	aoc.Run(func(r io.Reader) interface{} {
		solutions.PlayRps(r, handler)
		return score
	})
}
