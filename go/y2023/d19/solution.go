package main

import (
	"adventofcode/h"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type solution struct {
	input      []string
	ans1, ans2 int
}

func buildSolution(r io.Reader) *solution {
	lines, _ := h.LinesFromReader(r)
	return &solution{
		input: lines,
		ans1:  0,
		ans2:  0,
	}
}

// px{a<100:lmt,m>200:A,rfg}
type rule struct {
	field string
	op    byte
	value int
	next  string
}

type workflow struct {
	name     string
	rules    []rule
	defaultV string
}

type part struct {
	x, m, a, s int
}

func (s *solution) run1() {
	sep := 0
	workflows := make(map[string]workflow)
	for i, line := range s.input {
		if line == "" {
			sep = i
			break
		}
		wf := buildWorkflow(line)
		workflows[wf.name] = wf
	}

	for _, line := range s.input[sep+1:] {
		p := buildPart(line)
		res := getNext(p, workflows)
		if res == "A" {
			s.ans1 += p.a + p.m + p.s + p.x
		}
	}
}

func (p part) get(field string) int {
	switch field {
	case "x":
		return p.x
	case "m":
		return p.m
	case "s":
		return p.s
	case "a":
		return p.a
	}
	return 0
}

func (wf workflow) Next(p part) string {
	for _, r := range wf.rules {
		v := p.get(r.field)
		if r.op == '<' && v < r.value {
			return r.next
		}

		if r.op == '>' && v > r.value {
			return r.next
		}
	}
	return wf.defaultV
}

func getNext(p part, wf map[string]workflow) string {
	cur := "in"
	for cur != "A" && cur != "R" {
		cur = wf[cur].Next(p)
	}
	return cur
}

func buildPart(s string) part {
	strs := strings.Trim(s, "{}")
	p := part{}
	for part := range strings.SplitSeq(strs, ",") {
		key, value, _ := strings.Cut(part, "=")
		n, _ := strconv.Atoi(value)
		switch key {
		case "x":
			p.x = n
		case "m":
			p.m = n
		case "s":
			p.s = n
		case "a":
			p.a = n
		}
	}
	return p
}

func buildWorkflow(s string) workflow {
	name, body, _ := strings.Cut(s, "{")
	body = strings.TrimSuffix(body, "}")
	parts := strings.Split(body, ",")
	wf := workflow{
		name:     name,
		defaultV: parts[len(parts)-1],
	}
	for _, p := range parts[:len(parts)-1] {
		cond, next, _ := strings.Cut(p, ":")
		r := rule{
			field: cond[:1],
			op:    cond[1],
			next:  next,
		}
		r.value, _ = strconv.Atoi(cond[2:])
		wf.rules = append(wf.rules, r)
	}
	return wf
}

func (s *solution) run2() {
}

func (s *solution) res1() int {
	return s.ans1
}

func (s *solution) res2() int {
	return s.ans2
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res1()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	s.run2()
	return s.res2()
}

func main() {
	Input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("fail open input.txt %v", err)
	}
	start := time.Now()
	result := part1(Input)
	elapsed := time.Since(start)
	fmt.Printf("p1 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
	start = time.Now()
	result = part2(Input)
	elapsed = time.Since(start)
	fmt.Printf("p2 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
}
