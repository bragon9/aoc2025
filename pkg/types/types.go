package types

import "fmt"

func (s *Solution) Print() {
	fmt.Printf("~~~ Day %d ~~~\n", s.Day)
	fmt.Printf("Part 1: %v\n", s.Part1)
	fmt.Printf("Part 2: %v\n", s.Part2)
}

type Solution struct {
	Day   int
	Part1 any
	Part2 any
}
