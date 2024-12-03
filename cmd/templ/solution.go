package main

import (
	"adventofcode/preprocess"
	"fmt"
	"io"
	"os"
)

func main() {
	arg := os.Args[1]
	fmt.Println("Running part", arg)
	switch arg {
	case "1":
		res := Part1(os.Stdin)
		fmt.Println(res)
	case "2":
		res := Part2(os.Stdin)
		fmt.Println(res)
	}
}

func Part1(r io.Reader) int {
	_, _ = readLists(r)
	return 0
}

func Part2(r io.Reader) int {
	_, _ = readLists(r)
	return 0
}

func readLists(r io.Reader) ([][]int, error) {
	_, err := preprocess.LinesFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	return nil, err
}
