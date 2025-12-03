package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Telshin/advent-of-code-2025/util"
)

func Run() {
	// Get input text
	lines := util.ReadLines("day02/input.txt")

	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}

func solvePart1(lines []string) int {
	var total int = 0

	for _, line := range lines {
		ranges := strings.Split(line, ",")
		for _, pairs := range ranges {
			pair := strings.Split(pairs, "-")
			beginning, _ := strconv.Atoi(pair[0])
			end, _ := strconv.Atoi(pair[1])
			for i := beginning; i <= end; i++ {
				istr := strconv.Itoa(i)
				length := len(istr)

				if length%2 != 0 {
					continue
				}

				mid := length / 2
				if istr[:mid] == istr[mid:] {
					total += i
				}
			}
		}
	}

	return total
}

func solvePart2(lines []string) int {
	var total int = 0

	for _, line := range lines {
		ranges := strings.Split(line, ",")
		for _, pairs := range ranges {
			pair := strings.Split(pairs, "-")
			beginning, _ := strconv.Atoi(pair[0])
			end, _ := strconv.Atoi(pair[1])
			for i := beginning; i <= end; i++ {
				istr := strconv.Itoa(i)
				if checkForRepeatingCharacters(istr) {
					total += i
				}
			}
		}
	}

	return total
}

func checkForRepeatingCharacters(s string) bool {
	length := len(s)
	for x := 1; x <= length/2; x++ {
		if length%x == 0 {
			pattern := s[:x]
			numRepeats := length / x
			if s == strings.Repeat(pattern, numRepeats) {
				return true
			}
		}
	}
	return false
}
