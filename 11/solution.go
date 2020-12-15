package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
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

	fmt.Println(problemOne(values))
	fmt.Println(problemTwo(values))
}


func problemOne(values []int) int {
	return 0
}

func problemTwo(values []int) int {
	return 0
}