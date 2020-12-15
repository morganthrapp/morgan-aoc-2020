package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type action struct {
	action string
	value  int
}

type location struct {
	xPos    int
	yPos    int
	heading string
}

func (loc *location) move(act action) {
	switch act.action {
	case "N":
		loc.yPos = loc.yPos + act.value
	case "E":
		loc.xPos = loc.xPos + act.value
	case "S":
		loc.yPos = loc.yPos - act.value
	case "W":
		loc.xPos = loc.xPos - act.value
	case "R":
		for i := 0; i < (act.value / 90); i++ {
			loc.heading = turnRight(loc.heading)
		}
	case "L":
		for i := 0; i < (act.value / 90); i++ {
			loc.heading = turnLeft(loc.heading)
		}
	case "F":
		loc.move(action{
			value:  act.value,
			action: loc.heading,
		})
	}
}

func turnRight(current string) string {
	switch current {
	case "N":
		return "E"
	case "E":
		return "S"
	case "S":
		return "W"
	case "W":
		return "N"
	}
	panic(current)
}

func turnLeft(current string) string {
	switch current {
	case "N":
		return "W"
	case "W":
		return "S"
	case "S":
		return "E"
	case "E":
		return "N"
	}
	panic(current)
}

func (loc *location) turnRight() {
	switch loc.heading {
	case "E":
		newYPos := loc.xPos * -1
		loc.xPos = loc.yPos
		loc.yPos = newYPos
		loc.heading = "S"
	case "S":
		newYPos := loc.xPos * -1
		loc.xPos = loc.yPos
		loc.yPos = newYPos
		loc.heading = "W"
	case "W":
		newYPos := loc.xPos * -1
		loc.xPos = loc.yPos
		loc.yPos = newYPos
		loc.heading = "N"
	case "N":
		newYPos := loc.xPos * -1
		loc.xPos = loc.yPos
		loc.yPos = newYPos
		loc.heading = "E"
	default:
		panic(loc.heading)
	}
}

func (loc *location) turnLeft() {
	for i := 0; i < 3; i++ {
		loc.turnRight()
	}
}

func main() {
	rawData, err := ioutil.ReadFile("input")
	check(err)
	data := strings.Split(string(rawData), "\n")

	values := make([]action, len(data))
	for i, d := range data {
		actionType := string(d[0])
		value, _ := strconv.Atoi(d[1:])
		values[i] = action{
			actionType,
			value,
		}
	}

	fmt.Println(problemOne(values))
	fmt.Println(problemTwo(values))
}

func problemOne(actions []action) int {
	loc := location{
		heading: "E",
	}
	for _, action := range actions {
		loc.move(action)
	}
	return int(math.Abs(float64(loc.xPos)) + math.Abs(float64(loc.yPos)))
}

func problemTwo(actions []action) int {
	waypointLoc := location{
		heading: "E",
		yPos:    1,
		xPos:    10,
	}
	shipLoc := location{
		heading: "E",
	}
	for _, action := range actions {
		switch action.action {
		case "F":
			shipLoc.xPos = shipLoc.xPos + (waypointLoc.xPos * action.value)
			shipLoc.yPos = shipLoc.yPos + (waypointLoc.yPos * action.value)
		case "R":
			for i := 0; i < (action.value / 90); i++ {
				waypointLoc.turnRight()
			}
		case "L":
			for i := 0; i < (action.value / 90); i++ {
				waypointLoc.turnLeft()
			}
		default:
			waypointLoc.move(action)
		}
	}
	return int(math.Abs(float64(shipLoc.xPos)) + math.Abs(float64(shipLoc.yPos)))
}
