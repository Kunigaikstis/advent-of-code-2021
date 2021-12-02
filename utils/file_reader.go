package utils

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadNumsFromFile(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	r := bufio.NewReader(f)

	var done bool
	var nums []int

	for !done {
		input, err := r.ReadString('\n')
		if errors.Is(io.EOF, err) {
			done = true
		}

		num, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			return nil, err
		}

		nums = append(nums, num)
	}

	return nums, err
}
