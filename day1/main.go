package main

import (
	"fmt"
	"github.com/Kunigaikstis/advent-of-code-2021/utils"
	"math"
	"sync/atomic"
)

func main() {
	nums, err := utils.ReadNumsFromFile("day1/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(nums))
	fmt.Println(part2(nums))
}

func part1(nums []int) int32 {
	var count int32
	prev := math.MaxInt
	for _, num := range nums {
		if num > prev {
			atomic.AddInt32(&count, 1)
		}

		prev = num
	}

	return count
}

// a + b + c < b + c + d == a < d
func part2(nums []int) int32 {
	var count int32
	maxLength := len(nums) - 1
	for i := range nums {
		if i+3 > maxLength {
			continue
		}

		if nums[i+3] > nums[i] {
			atomic.AddInt32(&count, 1)
		}
	}

	return count
}