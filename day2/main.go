package main

import (
	"fmt"
	"github.com/Kunigaikstis/advent-of-code-2021/utils"
	"strconv"
	"strings"
)

func main() {
	inputs, err := utils.ReadStringsFromFile("day2/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(puzzle1(inputs))
	fmt.Println(puzzle2(inputs))
}

func puzzle1(inputs []string) int {
	hPos, depth := 0, 0

	for _, input := range inputs {
		command := strings.Split(input, " ")
		action := command[0]
		val, _ := strconv.Atoi(command[1])

		switch action {
		case "forward":
			hPos += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}

	return hPos * depth
}

func puzzle2(inputs []string) int {
	hPos, depth, aim := 0, 0, 0

	for _, input := range inputs {
		command := strings.Split(input, " ")
		action := command[0]
		val, _ := strconv.Atoi(command[1])

		switch action {
		case "forward":
			hPos += val
			depth += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}

	return hPos * depth
}

