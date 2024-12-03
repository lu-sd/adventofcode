package preprocess

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// return list of lines
func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := strings.TrimSpace(s.Text()) // Trim leading and trailing whitespace
		if line == "" {                     // Skip empty lines
			continue
		}
		lines = append(lines, line)
	}
	if s.Err() != nil {
		return nil, fmt.Errorf("failed to scan reader: %w", s.Err())
	}

	return lines, nil
}

func StrToInt(line string) ([]int, error) {
	wordlist := strings.Fields(line)
	ints := make([]int, len(wordlist))
	for i, s := range wordlist {
		val, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("failed to parse int %w", err)
		}
		ints[i] = val
	}
	return ints, nil
}
