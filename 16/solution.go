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

func any(value int, preds []func(int) bool) bool {
	for _, p := range preds {
		if p(value) {
			return true
		}
	}
	return false
}

func all(values []int, predicates []func(int) bool) bool {
	for _, v := range values {
		for _, p := range predicates {
			if !p(v) {
				return false
			}
		}
	}
	return true
}

func main() {
	rawData, err := ioutil.ReadFile("input")
	check(err)
	data := strings.Split(string(rawData), "\n")

	ruleRegex, _ := regexp.Compile(`(\w+): (\d+)-(\d+) or (\d+)-(\d+)`)
	rules := map[string][][2]int{}
	tickets := [][]int{}
	myTicket := []int{}
	mode := "rules"
	for _, l := range data {
		if l == "" {
			continue
		}
		if l == "your ticket:" {
			mode = "yt"
			continue
		}
		if l == "nearby tickets:" {
			mode = "nt"
			continue
		}
		switch mode {
		case "rules":
			ruleMatch := ruleRegex.FindStringSubmatch(l)
			firstStart, _ := strconv.Atoi(ruleMatch[2])
			firstEnd, _ := strconv.Atoi(ruleMatch[3])
			secondStart, _ := strconv.Atoi(ruleMatch[4])
			secondEnd, _ := strconv.Atoi(ruleMatch[5])
			rules[ruleMatch[1]] = [][2]int{{firstStart, firstEnd}, {secondStart, secondEnd}}
		case "nt":
			splitTicket := strings.Split(l, ",")
			ticket := make([]int, len(splitTicket))
			for i, t := range splitTicket {
				num, _ := strconv.Atoi(t)
				ticket[i] = num
			}
			tickets = append(tickets, ticket)
		case "yt":
			splitTicket := strings.Split(l, ",")
			myTicket = make([]int, len(splitTicket))
			for i, t := range splitTicket {
				num, _ := strconv.Atoi(t)
				myTicket[i] = num
			}
		}
	}

	fmt.Println(problemOne(rules, tickets))
	fmt.Println(problemTwo(rules, tickets, myTicket))
}

func problemOne(rules map[string][][2]int, tickets [][]int) int {
	ruleCheckers := make([]func(int) bool, len(rules)*2)
	i := 0
	for _, rule := range rules {
		for _, r := range rule {
			r := r
			f := func(x int) bool { return x >= r[0] && x <= r[1] }
			ruleCheckers[i] = f
			i++
		}
	}

	invalidTotal := 0
	for _, ticket := range tickets {
		for _, field := range ticket {
			if !any(field, ruleCheckers) {
				invalidTotal = invalidTotal + field
			}
		}
	}
	return invalidTotal
}

func problemTwo(rules map[string][][2]int, tickets [][]int, myTicket []int) int {
	ruleCheckers := map[string][]func(int) bool{}
	for ruleName, rule := range rules {
		for _, r := range rule {
			r := r
			f := func(x int) bool { return x >= r[0] && x <= r[1] }
			ruleCheckers[ruleName] = append(ruleCheckers[ruleName], f)
		}
	}

	ruleCheckerValues := make([]func(int) bool, len(ruleCheckers)*2)
	i := 0
	for _, v := range ruleCheckers {
		f1 := v[0]
		f2 := v[1]
		ruleCheckerValues[i] = func(x int) bool { return f1(x) || f2(x) }
		i++
	}

	validTickets := [][]int{}
	for _, ticket := range tickets {
		if all(ticket, ruleCheckerValues) {
			validTickets = append(validTickets, ticket)
		}
	}
	return 0
}
