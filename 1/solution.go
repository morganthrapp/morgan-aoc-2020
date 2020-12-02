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
	data := string(rawData)

	var values []int

	for _, x := range strings.Fields(data) {
		val, err := strconv.Atoi(x)
		check(err)
		values = append(values, val)
	}

	fmt.Println(problemOne(values))
	fmt.Println(problemTwo(values))
}

func problemOne(values []int) string {
	for i, a := range values {
		for _, b := range values[i+1:] {
			if a+b == 2020 {
				return fmt.Sprint(a * b)
			}
		}
	}
	return ""
}

func problemTwo(values []int) string {
	for i, a := range values {
		for x, b := range values[i+1:] {
			for _, c := range values[i+x+1:] {
				if a+b+c == 2020 {
					return fmt.Sprint(a * b * c)
				}
			}
		}
	}
	return ""
}
