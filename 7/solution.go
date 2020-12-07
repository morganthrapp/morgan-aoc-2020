package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Container is not exported, just to disambiguate with usage of container as a bare variable name
type Container struct {
	name       string
	containers []Container
}

func main() {
	rawData, err := ioutil.ReadFile("input")
	check(err)
	values := strings.Split(string(rawData), "\n")

	containsRegex, _ := regexp.Compile(`(\d+) (\w+ \w+) bags?`)

	containers := map[string]Container{}
	for _, value := range values {
		parts := strings.Split(value, " bags contain ")
		baseName := parts[0]
		subContainers := strings.Split(parts[1], ",")
		container := Container{
			baseName,
			[]Container{},
		}
		for _, sub := range subContainers {
			scMatch := containsRegex.FindStringSubmatch(sub)
			if scMatch == nil {
				continue
			}
			scNum, _ := strconv.Atoi(scMatch[1])
			subContainer := Container{
				name: scMatch[2],
			}
			for i := 0; i < scNum; i++ {
				container.containers = append(container.containers, subContainer)
			}
		}
		containers[baseName] = container
	}

	for _, value := range containers {
		value.populate(containers)
	}

	fmt.Println(problemOne("shiny gold", containers))
	fmt.Println(problemTwo(containers["shiny gold"]))
}

func (c Container) populate(mapping map[string]Container) {
	for i, container := range c.containers {
		c.containers[i].containers = mapping[container.name].containers
		container.populate(mapping)
	}
}

func (c Container) contains(name string) bool {
	for _, container := range c.containers {
		if container.name == name || container.contains(name) {
			return true
		}
	}
	return false
}

func (c Container) sum() int {
	count := len(c.containers)
	for _, container := range c.containers {
		count += container.sum()
	}
	return count
}

func problemOne(searchString string, containers map[string]Container) int {
	count := 0
	for _, value := range containers {
		if value.contains((searchString)) {
			count++
		}
	}
	return count
}

func problemTwo(container Container) int {
	return container.sum()
}
