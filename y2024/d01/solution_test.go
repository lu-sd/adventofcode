package main

import (
	"fmt"
	"strings"
	"testing"
)

const input = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func TestDistance(t *testing.T) {
	r := strings.NewReader(input)
	got, error := PartOne(r)
	fmt.Println(got, error)

	want := 11

	if got != want {
		t.Errorf("got error %v want %v, %v", got, want, error)
	}
}
