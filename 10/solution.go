package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func all(value interface{}, funcs []func(interface{}) bool) bool {
	for _, f := range funcs {
		if !f(value) {
			return false
		}
	}
	return true
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
	sort.Ints(values)
	values = append(values, values[len(values)-1]+3)

	fmt.Println(problemOne(values))
	fmt.Println(problemTwo(values))
}

func problemOne(values []int) int {
	oneDifs := 0
	threeDifs := 0
	previousNumber := 0
	for _, v := range values {
		if v-3 == previousNumber {
			threeDifs++
		}
		if v-1 == previousNumber {
			oneDifs++
		}
		previousNumber = v
	}
	return oneDifs * threeDifs
}

func problemTwo(values []int) int {
	counts := map[int]int{0: 1}
	for _, v := range values {
		counts[v] = counts[v-3] + counts[v-2] + counts[v-1]
	}
	return counts[values[len(values)-1]]
}
