package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func max(values []int) int {
	maxVal := values[0]
	for _, val := range values {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func main() {
	rawData, err := ioutil.ReadFile("input")
	check(err)
	values := strings.Split(string(rawData), "\n")

	fmt.Println(problemOne(values))
	fmt.Println(problemTwo(values))
}

func parseBoardingPass(pass string) int {
	passRegex, _ := regexp.Compile(`([FB]{7})([LR]{3})`)
	row := make([]int, 128)
	for i := range row {
		row[i] = i
	}
	column := make([]int, 8)
	for i := range column {
		column[i] = i
	}
	passMatches := passRegex.FindStringSubmatch(pass)
	rowSteps := passMatches[1]
	columnSteps := passMatches[2]
	for _, step := range rowSteps {
		newLen := len(row) / 2
		if string(step) == "F" {
			row = row[:newLen]
		} else {
			row = row[newLen:]
		}
	}
	for _, step := range columnSteps {
		newLen := len(column) / 2
		if string(step) == "L" {
			column = column[:newLen]
		} else {
			column = column[newLen:]
		}
	}
	seatID := (row[0] * 8) + column[0]
	return seatID
}

func problemOne(values []string) int {
	seatIds := make([]int, len(values))
	for i, pass := range values {
		seatIds[i] = parseBoardingPass(pass)
	}
	return max(seatIds)
}

func problemTwo(values []string) int {
	seatIds := make([]int, len(values))
	for i, pass := range values {
		seatIds[i] = parseBoardingPass(pass)
	}
	sort.Ints(seatIds)
	for i := 1; i < len(seatIds)-1; i++ {
		if !(seatIds[i-1]+1 == seatIds[i] && seatIds[i] == seatIds[i+1]-1) {
			return seatIds[i] + 1
		}
	}
	return 0
}
