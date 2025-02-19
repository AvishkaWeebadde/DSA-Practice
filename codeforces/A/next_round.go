package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

// Read a single integer
func nextInt() int {
	str, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(str))
	return n
}

// Read a space-separated list of integers
func nextInts() []int {
	str, _ := reader.ReadString('\n')
	strs := strings.Fields(str)
	nums := make([]int, len(strs))
	for i, s := range strs {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums
}

// Read a single string
func nextString() string {
	str, _ := reader.ReadString('\n')
	return strings.TrimSpace(str)
}

func main() {
	defer writer.Flush()

	// Read test cases
	inputs := nextInts()
	_, k := inputs[0], inputs[1]
	scores := nextInts()

	threshold := scores[k-1]
	count := 0

	for _, score := range scores {
		if score >= threshold && score > 0 {
			count++;
		}
	}

	// Output the result.
	fmt.Fprintln(writer, count)
}
