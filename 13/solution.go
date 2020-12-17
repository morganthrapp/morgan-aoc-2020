package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	rawData, err := ioutil.ReadFile("input")
	check(err)
	data := strings.Split(string(rawData), "\n")

	values := make([][]int, 2)
	earliestDeparture, _ := strconv.Atoi(data[0])
	values[0] = []int{earliestDeparture}
	values[1] = []int{}
	for _, b := range strings.Split(data[1], ",") {
		id, _ := strconv.Atoi(b)
		values[1] = append(values[1], id)
	}

	fmt.Println(problemOne(values))
	fmt.Println(problemTwo(values[1]))
}

func problemOne(timetable [][]int) int {
	target := timetable[0][0]
	earliestTime := target * 2
	earliestBusID := 0
	for _, b := range timetable[1] {
		if b == 0 {
			continue
		}
		bTime := b
		for bTime < target {
			bTime = bTime + b
			if bTime > target && bTime < earliestTime {
				earliestTime = bTime
				earliestBusID = b
			}
		}
	}
	return (earliestTime - target) * earliestBusID
}

func problemTwo(timetable []int) int {
	t, step := 0, timetable[0]

	for i, b := range timetable {
		if b == 0 || i == 0 {
			continue
		}
		for (t+i)%b != 0 {
			t = t + step
		}
		step = step * b
	}
	return t
}
