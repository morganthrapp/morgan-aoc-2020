package main

import (
	"fmt"
	"io/ioutil"
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

	questionare := ""
	values := []string{}
	for _, l := range data {
		if len(strings.TrimSpace(l)) == 0 {
			values = append(values, strings.TrimSpace(questionare))
			questionare = ""
			continue
		}
		questionare += "," + l
	}

	fmt.Println(problemOne(values))
	fmt.Println(problemTwo(values))
}

func problemOne(values []string) int {
	totalCount := 0
	for _, group := range values {
		answers := map[string]bool{}
		for _, person := range strings.Split(group, ",") {
			for _, answer := range person {
				answers[string(answer)] = true
			}
		}
		totalCount += len(answers)
	}
	return totalCount
}

func problemTwo(values []string) int {
	totalCount := 0
	for _, group := range values {
		answers := map[string]int{}
		people := strings.Split(group, ",")
		for _, person := range people {
			for _, answer := range person {
				answers[string(answer)]++
			}
		}
		for _, count := range answers {
			if count == len(people) - 1 {
				totalCount++
			}
		}
	}
	return totalCount
}
