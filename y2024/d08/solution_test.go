package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	testInput, err := os.Open("test1.txt")
	if err != nil {
		log.Fatalf("fail open test1.txt %v", err)
	}

	got := Part1(testInput)
	want := 14

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	testInput, err := os.Open("test1.txt")
	if err != nil {
		log.Fatalf("fail open test1.txt %v", err)
	}
	want := 34
	t.Run(fmt.Sprintf("want: %d", want), func(t *testing.T) {
		got := Part2(testInput)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
