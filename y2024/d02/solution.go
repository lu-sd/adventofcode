package main

import (
	"adventofcode/preprocess"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	arg := os.Args[1]
	fmt.Println("Running part", arg)
	switch arg {
	case "1":
		res := PartOne(os.Stdin)
		fmt.Println(res)
	case "2":
		res := PartTwo(os.Stdin)
		fmt.Println(res)
	}
}

func PartOne(r io.Reader) int {
	lines, err := readLists(r)
	if err != nil {
		return 0
	}
	var ans int
	for _, line := range lines {
		if isStrictM(line) && isValid(line) {
			ans++
		}
	}
	return ans
}

func PartTwo(r io.Reader) int {
	lines, err := readLists(r)
	if err != nil {
		return 0
	}

	var ans int
	for _, line := range lines {
		if isStrictM(line) && isValid(line) {
			ans++
			continue
		}

		for i := 0; i < len(line); i++ {
			newLine := []int{}
			newLine = append(newLine, line[:i]...)
			newLine = append(newLine, line[i+1:]...) // Remove the i-th element
			if isStrictM(newLine) && isValid(newLine) {
				ans++
				break // Stop further checks for this line
			}
		}
	}

	return ans
}

func readLists(r io.Reader) ([][]int, error) {
	lines, err := preprocess.LinesFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	var result [][]int
	for _, line := range lines {
		nums, err := preprocess.StrToInt(line)
		if err != nil {
			return nil, err
		}
		result = append(result, nums)
	}
	return result, nil
}

func isStrictM(l []int) bool {
	if len(l) < 2 {
		return true
	}
	inc := true
	dec := true
	for i := 1; i < len(l); i++ {
		if l[i] > l[i-1] {
			dec = false
		}
		if l[i] < l[i-1] {
			inc = false
		}
	}
	return inc || dec
}

func isValid(l []int) bool {
	for i := 1; i < len(l); i++ {
		diff := math.Abs(float64(l[i] - l[i-1]))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}
