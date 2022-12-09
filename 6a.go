package main

import (
	"fmt"
	"io"

	"github.com/snewell/advent2022/internal/aoc"
	"github.com/snewell/advent2022/internal/solutions"
)

func main() {
	handler := func(packet string) {
		fmt.Printf("%v\n", solutions.FindSequenceStart(packet, 4))
	}
	aoc.Run(func(r io.Reader) interface{} {
		solutions.ReadPackets(r, handler)
		return 0
	})
}
