package solutions

import (
	"fmt"
	"io"
)

type ClearingAssignment struct {
	Low  int
	High int
}

func ReadClearings(reader io.Reader, handler func(ClearingAssignment, ClearingAssignment)) {
	for {
		var first, second ClearingAssignment
		_, err := fmt.Fscanf(reader, "%d-%d,%d-%d", &first.Low, &first.High, &second.Low, &second.High)
		if err == nil {
			handler(first, second)
		} else {
			break
		}
	}
}
