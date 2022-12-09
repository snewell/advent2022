package solutions

import (
	"bufio"
	"io"
)

func considerSequence(packet string) int {
	// fmt.Printf("considering %v\n", packet)
	l := len(packet)
	for i := 0; i < (l - 1); i++ {
		for j := i + 1; j < l; j++ {
			if packet[i] == packet[j] {
				return i + 1
			}
		}
	}
	return 0
}

func FindSequenceStart(packet string, length int) int {
	index := 0
	for {
		offset := considerSequence(packet[index : index+length])
		if offset == 0 {
			return index + length
		}
		index += offset
	}
	return -1
}

func ReadPackets(reader io.Reader, handler func(string)) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		handler(text)
	}
}
