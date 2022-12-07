package solutions

import (
	"bufio"
	"fmt"
	"io"
)

type RpsSign int

type RpsResult int

const (
	RockSign     RpsSign = 1
	PaperSign    RpsSign = 2
	ScissorsSign RpsSign = 3

	RpsWin  RpsResult = 6
	RpsDraw RpsResult = 3
	RpsLoss RpsResult = 0
)

func CheckRpsRound(opponent RpsSign, you RpsSign) RpsResult {
	if opponent == you {
		return RpsDraw
	}
	if opponent > you {
		if opponent == ScissorsSign && you == RockSign {
			return RpsWin
		}
		return RpsLoss
	}
	if opponent == RockSign && you == ScissorsSign {
		return RpsLoss
	}
	return RpsWin
}

func mapOpponentSign(key byte) RpsSign {
	switch key {
	case 'A':
		return RockSign

	case 'B':
		return PaperSign

	case 'C':
		return ScissorsSign
	}
	panicMsg, _ := fmt.Printf("Unexpected sign: %v", key)
	panic(panicMsg)
	return RockSign
}

func mapYourSign(key byte) RpsSign {
	switch key {
	case 'X':
		return RockSign

	case 'Y':
		return PaperSign

	case 'Z':
		return ScissorsSign
	}
	panicMsg, _ := fmt.Printf("Unexpected sign: %v", key)
	panic(panicMsg)
	return RockSign
}

func mapResult(key byte) RpsResult {
	switch key {
	case 'X':
		return RpsLoss

	case 'Y':
		return RpsDraw

	case 'Z':
		return RpsWin
	}
	panicMsg, _ := fmt.Printf("Unexpected result: %v", key)
	panic(panicMsg)
	return RpsLoss
}

func PlayRps(reader io.Reader, handler func(RpsSign, RpsSign)) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		opponent := mapOpponentSign(text[0])
		you := mapYourSign(text[2])
		handler(opponent, you)
	}
}

func ExpectRps(reader io.Reader, handler func(RpsSign, RpsResult)) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		opponent := mapOpponentSign(text[0])
		you := mapResult(text[2])
		handler(opponent, you)
	}
}
