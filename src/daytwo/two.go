package main

import (
	"errors"
	"fmt"
)

type Program []int

func (p *Program) execute() {
	i := 0
	for i < len(*p) {
		instruction := (*p)[i]
		hop, err := p.do(instruction, i)
		if err != nil {
			fmt.Println("oops instruction not valid!")
			return
		}
		i += hop
	}
}

func (p *Program) do(instruction, slot int) (hop int, err error) {
	if instruction == 1 {
		// addition
		(*p)[(*p)[slot+3]] = (*p)[(*p)[slot+1]] + (*p)[(*p)[slot+2]]
		return 4, nil
	} else if instruction == 2 {
		// mult
		(*p)[(*p)[slot+3]] = (*p)[(*p)[slot+1]] * (*p)[(*p)[slot+2]]
		return 4, nil
	} else if instruction == 99 {
		return len(*p) + 1, nil
	}
	return 0, errors.New("no valid instruction found! instruction/slot:" + string(instruction) + " " + string(slot))
}

func partOne() {
	p := &Program{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 5, 19, 23, 2, 6, 23, 27, 1, 27, 5, 31, 2, 9, 31, 35, 1, 5, 35, 39, 2, 6, 39, 43, 2, 6, 43, 47, 1, 5, 47, 51, 2, 9, 51, 55, 1, 5, 55, 59, 1, 10, 59, 63, 1, 63, 6, 67, 1, 9, 67, 71, 1, 71, 6, 75, 1, 75, 13, 79, 2, 79, 13, 83, 2, 9, 83, 87, 1, 87, 5, 91, 1, 9, 91, 95, 2, 10, 95, 99, 1, 5, 99, 103, 1, 103, 9, 107, 1, 13, 107, 111, 2, 111, 10, 115, 1, 115, 5, 119, 2, 13, 119, 123, 1, 9, 123, 127, 1, 5, 127, 131, 2, 131, 6, 135, 1, 135, 5, 139, 1, 139, 6, 143, 1, 143, 6, 147, 1, 2, 147, 151, 1, 151, 5, 0, 99, 2, 14, 0, 0}
	p.execute()
	fmt.Println("state at pos 0", (*p)[0])
}

func calculate(noun, verb int) int {
	p := &Program{1, noun, verb, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 5, 19, 23, 2, 6, 23, 27, 1, 27, 5, 31, 2, 9, 31, 35, 1, 5, 35, 39, 2, 6, 39, 43, 2, 6, 43, 47, 1, 5, 47, 51, 2, 9, 51, 55, 1, 5, 55, 59, 1, 10, 59, 63, 1, 63, 6, 67, 1, 9, 67, 71, 1, 71, 6, 75, 1, 75, 13, 79, 2, 79, 13, 83, 2, 9, 83, 87, 1, 87, 5, 91, 1, 9, 91, 95, 2, 10, 95, 99, 1, 5, 99, 103, 1, 103, 9, 107, 1, 13, 107, 111, 2, 111, 10, 115, 1, 115, 5, 119, 2, 13, 119, 123, 1, 9, 123, 127, 1, 5, 127, 131, 2, 131, 6, 135, 1, 135, 5, 139, 1, 139, 6, 143, 1, 143, 6, 147, 1, 2, 147, 151, 1, 151, 5, 0, 99, 2, 14, 0, 0}
	p.execute()
	return (*p)[0]
}

func partTwo() {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			value := calculate(i, j)
			if value == 19690720 {
				// in case the solution is not unique
				// I don't return early
				solution := i*100 + j
				fmt.Printf("noun %v, verb %v, solution %v", i, j, solution)
			}
		}
	}
}

func main() {
	partOne()
	partTwo()
}
