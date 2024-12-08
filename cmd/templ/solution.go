package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)

type solution struct {
	ans int
}

func (s *solution) run1() {
}

func (s *solution) run2() {
}

func (s *solution) res() int {
	return 0
}

func buildSolution(r io.Reader) solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}

	return solution{}
}

func part1(r io.Reader) int {
	solution := buildSolution(r)
	solution.run1()
	return solution.res()
}

func part2(r io.Reader) int {
	solution := buildSolution(r)
	solution.run2()
	return solution.res()
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
