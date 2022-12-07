package solutions

import (
	"bufio"
	"io"
	"strconv"
)

func CountCalories(reader io.Reader, handler func(int)) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	calories := 0
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) != 0 {
			current, err := strconv.Atoi(text)
			if err != nil {
				panic(err)
			}
			calories += current
		} else {
			handler(calories)
			calories = 0
		}
	}
	if calories != 0 {
		handler(calories)
	}
}
