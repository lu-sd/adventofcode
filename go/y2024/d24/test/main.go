package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	input := strings.ReplaceAll(string(content), "\r\n", "\n")
	res, _ := SolutionForPart1(input)
	fmt.Println("paurt1", res)
	res2, _ := SolutionForPart2(input)
	fmt.Println("paurt1", res2)
}

type adjacencyList map[string][]string

func SolutionForPart1(input string) (int, error) {
	wireValues, wireDependencies := parseInput(input)
	return calculatePartOne(wireValues, wireDependencies), nil
}

func SolutionForPart2(input string) (string, error) {
	_, wireDependencies := parseInput(input)
	return calculatePartTwo(wireDependencies), nil
}

type wireDependency struct {
	wire1, wire2 string
	operation    string
}

func parseInput(input string) (map[string]int8, map[string]wireDependency) {
	instructionRegex := regexp.MustCompile(`([a-z0-9]*) ([A-Z]*) ([a-z0-9]*) -> ([a-z0-9]*)`)
	wireValueRegex := regexp.MustCompile(`([a-zA-Z0-9]*): ([0-9])`)

	scanner := bufio.NewScanner(strings.NewReader(input))

	wireValues := make(map[string]int8)
	wireDependencies := make(map[string]wireDependency)

	for scanner.Scan() && scanner.Text() != "" {
		matches := wireValueRegex.FindStringSubmatch(scanner.Text())
		wire := matches[1]
		value := int8(matches[2][0] - '0')
		wireValues[wire] = value
	}

	for scanner.Scan() {
		matches := instructionRegex.FindStringSubmatch(scanner.Text())
		wire := matches[4]
		operation := matches[2]
		wire1, wire2 := matches[1], matches[3]

		wireDependencies[wire] = wireDependency{
			wire1:     wire1,
			wire2:     wire2,
			operation: operation,
		}
	}

	return wireValues, wireDependencies
}

func calculatePartOne(wireValues map[string]int8, wireDependencies map[string]wireDependency) (result int) {
	var resolveWireValue func(string) int8

	resolveWireValue = func(currentWire string) int8 {
		if value, exists := wireValues[currentWire]; exists {
			return value
		}

		dependency := wireDependencies[currentWire]
		value1 := resolveWireValue(dependency.wire1)
		value2 := resolveWireValue(dependency.wire2)

		switch dependency.operation {
		case "XOR":
			wireValues[currentWire] = value1 ^ value2
		case "AND":
			wireValues[currentWire] = value1 & value2
		case "OR":
			wireValues[currentWire] = value1 | value2
		}

		return wireValues[currentWire]
	}

	for wire := range wireDependencies {
		resolveWireValue(wire)
	}

	for wire, value := range wireValues {
		if wire[0] == 'z' {
			position, _ := strconv.Atoi(wire[1:])
			result |= int(value) << position
		}
	}

	return
}

func calculatePartTwo(wireDependencies map[string]wireDependency) (result string) {
	validWires := make(map[string]bool)

	for wire, dependency := range wireDependencies {
		if wire[0] == 'z' {
			position, _ := strconv.Atoi(wire[1:])
			if dependency.operation != "XOR" && position != 45 {
				validWires[wire] = true
			}
		} else if !isWireXOrY(dependency.wire1) && !isWireXOrY(dependency.wire2) && dependency.wire1[0] != dependency.wire2[0] && dependency.operation == "XOR" {
			validWires[wire] = true
		}

		if dependency.operation == "XOR" && isWireXOrY(dependency.wire1) && isWireXOrY(dependency.wire2) && dependency.wire1[0] != dependency.wire2[0] {
			isValid := false
			for _, dep := range wireDependencies {
				if dep.operation == "XOR" && (dep.wire1 == wire || dep.wire2 == wire) {
					isValid = true
				}
			}
			if !isValid {
				validWires[wire] = true
			}
		}

		if dependency.operation == "AND" && isWireXOrY(dependency.wire1) && isWireXOrY(dependency.wire2) && dependency.wire1[0] != dependency.wire2[0] {
			isValid := false
			for _, dep := range wireDependencies {
				if dep.operation == "OR" && (dep.wire1 == wire || dep.wire2 == wire) {
					isValid = true
				}
			}
			if !isValid {
				validWires[wire] = true
			}
		}
	}
	sortedWires := slices.Collect(maps.Keys(validWires))
	slices.Sort(sortedWires)

	for _, wire := range sortedWires {
		result += wire + ","
	}

	return result[:len(result)-1]
}

func isWireXOrY(wire string) bool {
	position, _ := strconv.Atoi(wire[1:])
	return (wire[0] == 'x' || wire[0] == 'y') && position != 0
}
