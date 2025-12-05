package solutions

import (
	"aoc2025/pkg/parser"
	"aoc2025/pkg/types"
	"strconv"
	"strings"
)

func isValidID(id int) bool {
	stringID := strconv.Itoa(id)

	if len(stringID)%2 != 0 {
		return true
	}

	left := stringID[:len(stringID)/2]
	right := stringID[len(stringID)/2:]

	return left != right
}

func getInvalidIDsForRange(start, stop int) []int {
	var invalidIDs []int

	for i := start; i <= stop; i++ {
		if !isValidID(i) {
			invalidIDs = append(invalidIDs, i)
		}
	}

	return invalidIDs
}

func getSumOfInvalidIDs(lines []string) int {
	var sum int
	for _, line := range lines {
		ranges := strings.Split(line, "-")
		start, _ := strconv.Atoi(ranges[0])
		stop, _ := strconv.Atoi(ranges[1])
		invalidIDs := getInvalidIDsForRange(start, stop)

		for _, id := range invalidIDs {
			sum += id
		}
	}

	return sum
}

func day2p1() any {
	line, err := parser.ReadFile("/Users/brandon.spagon/git/aoc2025/pkg/inputs/day2.txt")
	if err != nil {
		return nil
	}

	lines := strings.Split(line, ",")
	return getSumOfInvalidIDs(lines)
}

func day2p2() any {
	return 1
}

func SolveDay2() types.Solution {
	return types.Solution{
		Day:   2,
		Part1: day2p1(),
		Part2: day2p2(),
	}
}
