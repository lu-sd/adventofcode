package main

import (
	"fmt"
	"strings"
	"testing"
)

const input1 = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func TestDistance(t *testing.T) {
	r := strings.NewReader(input1)
	got, error := PartOne(r)
	fmt.Println(got, error)

	want := 11

	if got != want {
		t.Errorf("got error %v want %v, %v", got, want, error)
	}
}

func TestPart2(t *testing.T) {
	r := strings.NewReader(input1)
	got, error := PartTwo(r)
	fmt.Println(got, error)

	want := 31

	if got != want {
		t.Errorf("got error %v want %v, %v", got, want, error)
	}
}
