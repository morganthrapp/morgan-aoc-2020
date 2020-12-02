package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func must(i int, e error) int {
	check(e)
	return i
}

func main() {
	rawData, err := ioutil.ReadFile("input")
	check(err)
	data := strings.Split(string(rawData), "\n")
	
	var values [][]string

	r, _ := regexp.Compile(`(\d+)-(\d+) (\w): (\w+)`)

	for _, p := range data {
		match := r.FindStringSubmatch(p)
		values = append(values, match[1:])
	}

	fmt.Println(problemOne(values))
	fmt.Println(problemTwo(values))
}

func problemOne(values [][]string) int {
	count := 0
	for _, x := range values {
		start := must(strconv.Atoi(x[0]))
		end := must(strconv.Atoi(x[1]))
		chr := x[2]
		pw := x[3]
		pwCount := strings.Count(pw, chr)
		if start <= pwCount && pwCount <= end {
			count++
		}
	}	
	return count
}

func problemTwo(values [][]string) int {
	count := 0
	for _, x := range values {
		firstPos := must(strconv.Atoi(x[0]))
		secondPos := must(strconv.Atoi(x[1]))
		chr := x[2]
		pw := x[3]
		pwCount := 0
		if string(pw[firstPos-1]) == chr {
			pwCount++
		}
		if string(pw[secondPos-1]) == chr {
			pwCount++
		}
		if pwCount == 1 {
			count++
		}
	}	
	return count
}