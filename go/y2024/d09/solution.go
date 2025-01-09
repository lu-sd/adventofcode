package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type solution struct {
	input    []byte
	intSlice []int
	ans      int
}

func (s *solution) run1() {
	s.buildIntSlice()
	s.reorderSlice()
}

func (s *solution) buildIntSlice() {
	fId := 0

	for i, b := range s.input {
		if b == '\n' {
			break
		}

		length := int(b - '0')
		if i%2 == 0 {
			s.intSlice = append(s.intSlice, newSlice(fId, length)...)
			fId++
		} else {
			s.intSlice = append(s.intSlice, newSlice(-1, length)...)
		}
	}
}

func newSlice(fill, length int) []int {
	s := make([]int, length)
	for i := range s {
		s[i] = fill
	}
	return s
}

func (s *solution) reorderSlice() {
	// 0-1-1111-1-1-122
	// 022111
	l, r := 0, len(s.intSlice)-1
	for l < r {
		if s.intSlice[l] != -1 {
			l++
			continue
		}

		if s.intSlice[r] == -1 {
			r--
			continue
		}
		// swap
		s.intSlice[l], s.intSlice[r] = s.intSlice[r], s.intSlice[l]
		l++
		r--
	}
}

func (s *solution) run2() {
	s.buildIntSlice()
	s.reorderSlice2()
}

func (s *solution) reorderSlice2() {
	for fId := s.intSlice[len(s.intSlice)-1]; fId >= 0; fId-- {
		fStart, fLength := s.findFStartNLen(fId)
		lStart, find := s.findSpot(fLength, fStart)
		if find {
			// swap
			for i := 0; i < fLength; i++ {
				s.intSlice[fStart+i] = -1
				s.intSlice[lStart+i] = fId
			}
		}
	}
}

func (s *solution) findSpot(reqLenth, end int) (int, bool) {
	curlen, start := 0, -1
	for i := 0; i < end; i++ {
		if s.intSlice[i] == -1 {
			curlen++
			if start == -1 {
				start = i
			}
			if curlen == reqLenth {
				return start, true
			}
			continue
		}
		curlen, start = 0, -1
	}
	return -1, false
}

func (s *solution) findFStartNLen(fId int) (start, flen int) {
	for i, num := range s.intSlice {
		if num == fId {
			start = i
			break
		}
	}

	for i, num := range s.intSlice {
		if i < start {
			continue
		}
		if num != fId {
			break
		}
		flen++
	}
	return start, flen
}

func (s *solution) res() int {
	for i, num := range s.intSlice {
		if num == -1 {
			continue
		}
		s.ans += i * num
	}

	return s.ans
}

func buildSolution(r io.Reader) *solution {
	// lines, err := utils.LinesFromReader(r)
	line, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", line, err)
	}

	return &solution{
		input: line,
	}
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	s.run2()
	return s.res()
}

func main() {
	arg := os.Args[1]
	fmt.Println("Running part", arg)
	switch arg {
	case "1":
		fmt.Println("p1 res 🙆-> ", part1(os.Stdin))
	case "2":
		fmt.Println("p2 res 🙆-> ", part2(os.Stdin))
	}
}
