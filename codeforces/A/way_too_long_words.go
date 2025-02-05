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
	t := nextInt()
	for i := 0; i < t; i++ {
		s := nextString()
		if len(s) <= 10 {
			fmt.Fprintln(writer, s)
		} else {
			fmt.Fprintf(writer, "%c%d%c\n", s[0], len(s)-2, s[len(s)-1])
		}
	}
}
