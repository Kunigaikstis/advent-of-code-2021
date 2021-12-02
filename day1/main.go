package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
)

func main() {
	f, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	r := bufio.NewReader(f)

	var count int32
	prev := math.MaxInt
	var done bool
	var nums []int

	for !done {
		input, err := r.ReadString('\n')
		if errors.Is(io.EOF, err) {
			done = true
		}

		num, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}

	for _, num := range nums {
		if num > prev {
			atomic.AddInt32(&count, 1)
		}

		prev = num
	}

	fmt.Println(count)
}