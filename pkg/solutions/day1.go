package solutions

import (
	"aoc2025/pkg/parser"
	"aoc2025/pkg/types"
	"strconv"
)

type instruction struct {
	direction string
	amount    int
}

func parseLinesToInstructions(lines []string) []instruction {
	instructions := make([]instruction, 0, len(lines))
	for _, line := range lines {
		var dir string
		var amt int

		dir = string(line[0])
		amt, _ = strconv.Atoi(line[1:])

		instructions = append(instructions, instruction{
			direction: dir,
			amount:    amt,
		})
	}

	return instructions
}

func day1p1() any {
	lines, err := parser.ReadFileLines("/Users/brandon.spagon/git/aoc2025/pkg/inputs/day1.txt")
	if err != nil {
		return nil
	}

	instructions := parseLinesToInstructions(lines)
	var password int
	curr := 50
	for _, instr := range instructions {
		if instr.direction == "L" {
			curr -= instr.amount
		}
		if instr.direction == "R" {
			curr += instr.amount
		}
		if curr%100 == 0 {
			password += 1
		}
	}

	return password
}

func day1p2() any {
	lines, err := parser.ReadFileLines("/Users/brandon.spagon/git/aoc2025/pkg/inputs/day1.txt")
	if err != nil {
		return nil
	}

	instructions := parseLinesToInstructions(lines)
	var password int
	curr := 50
	for _, instr := range instructions {
		dialValue := curr % 100
		if dialValue < 0 {
			dialValue += 100
		}

		password += instr.amount / 100
		remainder := instr.amount % 100
		if dialValue != 0 {
			if instr.direction == "L" {
				if remainder >= dialValue {
					password += 1
				}
			} else {
				if remainder >= (100 - dialValue) {
					password += 1
				}
			}
		}

		if instr.direction == "L" {
			curr -= instr.amount
		}
		if instr.direction == "R" {
			curr += instr.amount
		}
	}

	return password
}

func SolveDay1() types.Solution {
	return types.Solution{
		Day:   1,
		Part1: day1p1(),
		Part2: day1p2(),
	}
}
