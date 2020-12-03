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

func must(i int, e error) int {
	check(e)
	return i
}

func main() {
	rawData, err := ioutil.ReadFile("input")
	check(err)
	values := strings.Split(string(rawData), "\n")

	fmt.Println(problemOne(values))
	fmt.Println(problemTwo(values))
}

func problemOne(values []string) int {
	return countTrees(values, 3, 1)
}

func countTrees(values []string, xMovement int, yMovement int) int {
	xLoc := 0
	yLoc := 0
	treeCount := 0
	lineLength := len(values[yLoc])
	for yLoc < len(values) - yMovement {
		yLoc = yLoc + yMovement
		xLoc = (xLoc + xMovement) % lineLength
		if string(values[yLoc][xLoc]) == "#" {
			treeCount++
		}
	}
	return treeCount
}

func problemTwo(values []string) int {
	curves := [][]int{ {1,1}, {3,1}, {5,1}, {7,1}, {1,2} }
	total := 1
	for _, c := range curves {
		x := c[0]
		y := c[1]
		total = total * countTrees(values, x, y)
	}
	return total
}
