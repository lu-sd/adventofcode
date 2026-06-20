package h

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func ByteSFromReader(r io.Reader) ([]byte, error) {
	line, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("error %w", err)
	}
	if seeker, ok := r.(io.Seeker); ok {
		seeker.Seek(0, 0)
	}

	return bytes.TrimSpace(line), nil
}

func IsDigitBool(v any) bool {
	_, ok := IsDigit(v)
	return ok
}

func IsDigit(v any) (int, bool) {
	var r rune

	switch x := v.(type) {
	case rune:
		r = x
	case byte:
		r = rune(x)
	default:
		return 0, false
	}

	if r >= '0' && r <= '9' {
		return int(r - '0'), true
	}
	return 0, false
}

// LinesFromReader return list of lines
func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if s.Err() != nil {
		return nil, fmt.Errorf("failed to scan reader: %w", s.Err())
	}

	if seeker, ok := r.(io.Seeker); ok {
		seeker.Seek(0, 0)
	}
	return lines, nil
}

// func IntsFromString(line string) ([]int, error) {
// 	wordlist := strings.Fields(line)
// 	ints := make([]int, len(wordlist))
// 	for i, s := range wordlist {
// 		val, err := strconv.Atoi(s)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to parse int %w", err)
// 		}
// 		ints[i] = val
// 	}
// 	return ints, nil
// }
