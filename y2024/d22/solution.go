package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func (s *solution) run1() {
	for _, line := range s.lines {
		item := uint64(utils.IntsFromString(line)[0])
		try := &secert{item}
		for range 2000 {
			try.next()
		}
		s.ans += int(try.val)
	}
}

func (s *solution) run2() {
	for _, line := range s.lines {
		item := uint64(utils.IntsFromString(line)[0])
		try := &secert{item}
		candidatas := map[seq]int{}
		var price, diff []int
		prev := 0
		for range 2000 {
			v := try
			lastbit := int(v.val % 10)
			price = append(price, lastbit)
			diff = append(diff, lastbit-prev)
			prev = lastbit
			try.next()
		}
		// fmt.Println(price)
		// fmt.Println(diff)
		// build seq
		for i := 1; i < len(price)-4; i++ {
			var curSeq seq
			for j := range 4 {
				curSeq[j] = diff[i+j]
			}
			if _, ok := candidatas[curSeq]; ok {
				continue
			}
			candidatas[curSeq] = price[i+3]
		}

		for k, v := range candidatas {
			s.prices[k] += v
		}
		// fmt.Printf("%v", candidatas)
	}

	for _, v := range s.prices {
		s.ans = max(s.ans, v)
	}
}

func (s *solution) res() int {
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}

	return &solution{
		lines:  lines,
		prices: map[seq]int{},
		ans:    0,
	}
}

type (
	seq      [4]int
	solution struct {
		lines  []string
		ans    int
		prices map[seq]int
	}
)

type secert struct {
	val uint64
}

func (s *secert) next() *secert {
	s.mix(-6).prune().mix(5).prune().mix(-11).prune()
	return s
}

func (s *secert) mix(off int) *secert {
	if off < 0 {
		s.val = s.val ^ s.val<<(-off)
	} else {
		s.val = s.val ^ s.val>>off
	}
	return s
}

func (s *secert) prune() *secert {
	s.val = s.val % 16777216
	return s
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
	Input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("fail open input.txt %v", err)
	}
	start := time.Now()
	result := part1(Input)
	elapsed := time.Since(start)
	fmt.Printf("p1 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
	start = time.Now()
	result = part2(Input)
	elapsed = time.Since(start)
	fmt.Printf("p2 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
}
