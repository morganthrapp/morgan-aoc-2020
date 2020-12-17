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

func parseMask(mask string) []int {
	mask = reverse(strings.TrimPrefix(mask, "mask = "))
	res := make([]int, len(mask))
	for i, m := range mask {
		if string(m) == "X" {
			res[i] = -1
		}
		if string(m) == "1" {
			res[i] = 1
		}
		if string(m) == "0" {
			res[i] = 0
		}
	}
	return res
}

func reverse(s string) string {
	res := ""
	for _, c := range s {
		res = string(c) + res
	}
	return res
}

func sum(vals map[int]int) int {
	res := 0
	for _, v := range vals {
		res = res + v
	}
	return res
}

func iToA(n int, sequence ...int) []int {
	if n != 0 {
		i := n % 10
		sequence = append(sequence, i)
		return iToA(n/10, sequence...)
	}
	return sequence
}

func main() {
	rawData, err := ioutil.ReadFile("input")
	check(err)
	data := strings.Split(string(rawData), "\n")

	mask := ""
	values := map[string][][2]int{}
	addressRegex, _ := regexp.Compile(`mem\[(\d+)\] = (\d+)`)
	maskRegex, _ := regexp.Compile(`mask = (.*)`)
	program := [][2]int{}
	for _, l := range data {
		maskMatch := maskRegex.FindStringSubmatch(l)
		if maskMatch != nil {
			if len(mask) > 0 {
				values[mask] = program
				program = [][2]int{}
			}
			mask = maskMatch[1]
		} else {
			match := addressRegex.FindStringSubmatch(l)
			adr, _ := strconv.Atoi(match[1])
			val, _ := strconv.Atoi(match[2])
			program = append(program, [2]int{adr, val})
		}
	}

	values[mask] = program

	fmt.Println(problemOne(values))
	fmt.Println(problemTwo(values))
}

func problemOne(values map[string][][2]int) int {
	memory := map[int]int{}
	for mask, instructions := range values {
		mask := parseMask(mask)
		for _, ins := range instructions {
			address := ins[0]
			value, _ := strconv.Atoi(strconv.FormatInt(int64(ins[1]), 2))
			parsedValue := iToA(value)
			res := make([]string, len(mask))
			for i := 0; i < len(mask); i++ {
				if mask[i] == -1 {
					if i < len(parsedValue) {
						res[i] = fmt.Sprint(parsedValue[i])
					} else {
						res[i] = "0"
					}
				} else {
					res[i] = fmt.Sprint(mask[i])
				}
			}
			finalValue, _ := strconv.ParseInt(reverse(strings.Join(res, "")), 2, len(mask))
			memory[address] = int(finalValue)
		}
	}
	return sum(memory)
}

func problemTwo(values map[string][][2]int) int {
	return 0
}
