package main

import (
	"fmt"
	"github.com/Kunigaikstis/advent-of-code-2021/utils"
	"strconv"
	"sync/atomic"
)

func main() {
	input, err := utils.ReadStringsFromFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(puzzle1(input))
}

func puzzle1(input []string) int {
	inputLength:=len(input[0])
	ones := make([]int32, inputLength, inputLength)
	for _, num := range input {
		for i, chr := range num {
			if chr == 49 {
				atomic.AddInt32(&ones[i], 1)
			}
		}
	}

	gammaStr, epsilonStr := "", ""
	totalInputs := int32(len(input))

	for _, col := range ones {
		if col > totalInputs/2 {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}

	fmt.Println(gammaStr)
	fmt.Println(epsilonStr)

	gamma, _ := strconv.ParseInt(gammaStr, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 64)

	return int(gamma * epsilon)
}