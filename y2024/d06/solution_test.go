package main

import (
	"strings"
	"testing"
)

const input1 = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestPart1(t *testing.T) {
	r := strings.NewReader(input1)
	got, _ := Part1(r)
	want := 41

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	r := strings.NewReader(input1)
	got, _ := Part2(r)
	want := 6

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
