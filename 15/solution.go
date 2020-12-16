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
	data := strings.Split(strings.Split(string(rawData), "\n")[0], ",")

	values := make([]int, len(data))
	for i, l := range data {
		num, _ := strconv.Atoi(l)
		values[i] = num
	}

	fmt.Println(problemOne(values, 2020))
	fmt.Println(problemTwo(values))
}

func problemOne(values []int, target int) int {
	t := len(values)
	numHistory := map[int]int{}
	for i, v := range values[:len(values)-1] {
		numHistory[v] = i + 1
	}
	curNum := values[len(values)-1]
	for t < target {
		lastTime := numHistory[curNum]
		nextNum := 0
		if lastTime != 0 {
			nextNum = t - lastTime
		}
		numHistory[curNum] = t
		curNum = nextNum
		t++
	}
	return curNum
}

func problemTwo(values []int) int {
	return problemOne(values, 30000000)
}
