package day01

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/Telshin/advent-of-code-2025/util"
)

func Run() {
	// Get input text
	lines := util.ReadLines("day01/input.txt")

	// Variables
	var dial int = 50

	fmt.Println("Part 1:", solvePart1(lines, dial))
	fmt.Println("Part 2:", solvePart2(lines, dial))
}

func solvePart1(lines []string, dial int) int {
	// Setup variables to keep track of everything
	var password int = 0
	var dialValue = dial

	for _, line := range lines {
		r := regexp.MustCompile(`(L|R)(\d+)`)
		match := r.FindStringSubmatch(line)
		num, _ := strconv.Atoi(match[2])

		switch match[1] {
		case "L":
			dialValue -= num
		case "R":
			dialValue += num
		}

		if dialValue%100 == 0 {
			password++
		}
	}

	return password
}

func solvePart2(lines []string, dial int) int {
	// Setup variables to keep track of everything
	var password int = 0
	var dialValue = dial

	for _, line := range lines {
		r := regexp.MustCompile(`(L|R)(\d+)`)
		match := r.FindStringSubmatch(line)
		num, _ := strconv.Atoi(match[2])

		// Calculate full rotations from the original distance
		fullRotations := int(math.Floor(float64(num) / 100))
		password += fullRotations
		num %= 100

		// Landing on zero
		if dialValue == 0 {
			password++
		}

		switch match[1] {
		case "L":
			// Partial left rotation passing zero
			if num > dialValue {
				if dialValue != 0 {
					password++
				}
				dialValue = (dialValue - num) + 100
			} else {
				dialValue = (dialValue - num)
			}
		case "R":
			// Partial right rotation passing zero
			if (dialValue + num) > 100 {
				password++
			}
			dialValue = (dialValue + num) % 100
		}
	}

	return password
}
