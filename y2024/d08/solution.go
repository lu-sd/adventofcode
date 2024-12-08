package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"os"
)

func main() {
	arg := os.Args[1]
	fmt.Println("Running part", arg)
	switch arg {
	case "1":
		res, err := Part1(os.Stdin)
		if err != nil {
			fmt.Println("p1 error ", err)
		}
		fmt.Println("p1 res 🙆-> ", res)
	case "2":
		res, err := Part2(os.Stdin)
		if err != nil {
			fmt.Println("p2 error ", err)
		}
		fmt.Println("p2 res 🙆-> ", res)
	}
}

func Part1(r io.Reader) (int, error) {
	lines, err := readLists(r)
	if err != nil {
		fmt.Println(lines)
		return 0, fmt.Errorf("error %w", err)
	}

	return 0, nil
}

func Part2(r io.Reader) (int, error) {
	lines, err := readLists(r)
	if err != nil {
		fmt.Println(lines)
		return 0, fmt.Errorf("error %w", err)
	}
	return 0, nil
}

func readLists(r io.Reader) ([]string, error) {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	return lines, err
}
