package solutions

import (
	"bufio"
	"io"
	"unicode"
)

func GetPriority(item rune) int {
	if unicode.IsUpper(item) {
		return int(item - 'A' + 27)
	}
	return int(item - 'a' + 1)
}

func Rucksack(reader io.Reader, handler func(string, string)) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		halfLength := len(text) / 2
		handler(text[:halfLength], text[halfLength:])
	}
}
