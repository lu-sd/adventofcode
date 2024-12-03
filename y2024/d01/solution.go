package main

import (
	"adventofcode/preprocess"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	res, err := PartOne(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func distanceBetweenLists(a, b []int) int {
	var distance int
	for i := range a {
		distance += abs(a[i] - b[i])
	}
	return distance
}

func PartOne(r io.Reader) (int, error) {
	list1, list2, err := readLists(r)
	if err != nil {
		return 0, err
	}
	slices.Sort(list1)
	slices.Sort(list2)

	distance := distanceBetweenLists(list1, list2)
	return distance, nil
}

func readLists(r io.Reader) ([]int, []int, error) {
	lines, err := preprocess.LinesFromReader(r)
	if err != nil {
		return nil, nil, fmt.Errorf("could not read input: %w", err)
	}

	var (
		col1 = make([]int, len(lines))
		col2 = make([]int, len(lines))
	)

	for i, line := range lines {
		n1, n2, err := LinesTo2Int(line)
		if err != nil {
			return nil, nil, err
		}
		col1[i] = n1
		col2[i] = n2
	}

	return col1, col2, nil
}

func LinesTo2Int(line string) (int, int, error) {
	nums := strings.Fields(line)
	res := []int{}
	for _, v := range nums {
		num, err := strconv.Atoi(v)
		if err == nil {
			res = append(res, num)
		}
	}
	if len(res) == 2 {
		return res[0], res[1], nil
	}

	return 0, 0, fmt.Errorf("ParseInt error")
}
