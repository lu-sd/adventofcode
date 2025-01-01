package main

import (
	"fmt"
	"io"
	"os"
)

type solution struct {
	lines []byte
}

func (s *solution) mul() int {
	res := 0
	line := string(s.lines)
	for i := 0; i < len(line)-4; i++ {
		if line[i:i+4] != "mul(" {
			continue
		}
		i += 4
		n1 := s.getNum(&i, line)
		if line[i] != ',' {
			continue
		}
		i++
		n2 := s.getNum(&i, line)
		if line[i] != ')' {
			continue
		}
		res += n1 * n2
	}
	return res
}

func (s *solution) mul2() int {
	line := string(s.lines)
	res := 0
	enable := true
	for i := 0; i < len(line)-7; i++ {
		if line[i:i+7] == "don't()" {
			enable = false
		}
		if line[i:i+4] == "do()" {
			enable = true
		}
		if line[i:i+4] == "mul(" {
			i += 4
			n1 := s.getNum(&i, line)
			if line[i] != ',' {
				continue
			}
			i++
			n2 := s.getNum(&i, line)
			if line[i] != ')' {
				continue
			}
			if enable {
				res += n1 * n2
			}
		}
	}
	return res
}

func (s *solution) getNum(idx *int, line string) int {
	num := 0
	for *idx < len(line) && line[*idx] <= '9' && line[*idx] >= '0' {
		num = num*10 + int(line[*idx]-'0')
		*idx++
	}
	return num
}

func buildSolution(r io.Reader) *solution {
	lines, err := io.ReadAll(r)
	if err != nil {
		return nil
	}

	return &solution{
		lines: lines,
	}
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	res := s.mul()
	return res
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	res := s.mul2()
	return res
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
