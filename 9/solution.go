package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(values []int) int {
	total := 0
	for _, n := range values {
		total += n
	}
	return total
}

func main() {
	rawData, err := ioutil.ReadFile("input")
	check(err)
	data := strings.Split(string(rawData), "\n")

	values := make([]int, len(data))
	for i, d := range data {
		num, _ := strconv.Atoi(d)
		values[i] = num
	}

	fmt.Println(problemOne(values))
	fmt.Println(problemTwo(values))
}

func problemOne(values []int) int {
	preambleLength := 25
	for i, n := range values {
		if i < preambleLength {
			continue
		}
		validNumbers := calculatePreamble(values[i-preambleLength : i])
		if !validNumbers[n] {
			return n
		}
	}
	return 0
}

func problemTwo(values []int) int {
	target := problemOne(values)
	current := 0
	result := 0
	start := 0
	end := 1
	for end < len(values) {
		slice := values[start:end]
		current = sum(slice)
		end++
		if current > target {
			start++
			end = start + 1
		}
		if current == target {
			sort.Ints(slice)
			result = slice[0] + slice[len(slice)-1]
			break
		}
	}
	return result
}

func calculatePreamble(values []int) map[int]bool {
	validNumbers := make(map[int]bool, len(values))
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values); j++ {
			if i == j {
				continue
			}
			validNumbers[values[i]+values[j]] = true
		}
	}
	return validNumbers
}
