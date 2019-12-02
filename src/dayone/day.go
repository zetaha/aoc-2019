package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func fuelCounterUpper(moduleW int) int {
	return int(math.Floor(float64(moduleW)/3.0)) - 2
}

func parseInput(filename string) (moduleWeights []int, err error) {
	F, err := os.Open(filename)
	if err != nil {
		fmt.Println("opening file returned an error", err)
		return moduleWeights, err
	}
	scanner := bufio.NewScanner(F)
	for scanner.Scan() {
		line := scanner.Text()
		input, _ := strconv.Atoi(line)
		moduleWeights = append(moduleWeights, input)
	}
	return moduleWeights, nil
}

func partOne() {
	moduleWeights, err := parseInput("1.txt")
	if err != nil {
		fmt.Println("error parsing the file", err)
		return
	}

	sum := 0
	for _, w := range moduleWeights {
		sum += fuelCounterUpper(w)
	}
	fmt.Println("sum of the fuel required for the modules", sum)
}

// total fuel accounts for the fuel of the fuel
func totalFuel(w int) int {
	sum := 0
	for w > 0 {
		fuel := fuelCounterUpper(w)
		w = fuel
		if fuel >= 0 {
			sum += fuel
		}
	}
	return sum
}

func partTwo() {
	moduleWeights, err := parseInput("1.txt")
	if err != nil {
		fmt.Println("error loading file of inputs", err)
		return
	}
	sum := 0
	for _, w := range moduleWeights {
		sum += totalFuel(w)
	}
	fmt.Println("total fuel accounting the weight of the fuel itself", sum)
}

func main() {
	partOne()
	partTwo()
}
