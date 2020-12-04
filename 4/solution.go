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

	passport := ""
	values := []string{}
	for _, l := range data {
		if len(strings.TrimSpace(l)) == 0 {
			values = append(values, passport)
			passport = ""
			continue
		}
		passport = passport + " " + l
	}

	fmt.Println(problemOne(values))
	fmt.Println(problemTwo(values))
}

func problemOne(values []string) int {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	requiredFieldFunctions := []func(interface{}) bool{}
	for _, field := range requiredFields {
		field := field
		f := func(passport interface{}) bool {
			return strings.Contains(passport.(string), field)
		}
		requiredFieldFunctions = append(requiredFieldFunctions, f)
	}

	validCount := 0
	for _, passport := range values {
		if all(passport, requiredFieldFunctions) {
			validCount++
		}
	}
	return validCount
}

func problemTwo(values []string) int {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	requiredFieldFunctions := []func(interface{}) bool{}

	for _, field := range requiredFields {
		field := field
		f := func(passport interface{}) bool {
			return passport.(map[string]string)[field] != ""
		}
		requiredFieldFunctions = append(requiredFieldFunctions, f)
	}

	heightRegex, _ := regexp.Compile(`^(\d{2,3})(cm|in)$`)
	hairColorRegex, _ := regexp.Compile(`^#[a-f0-9]{6}$`)
	passportIDRegex, _ := regexp.Compile(`^\d{9}$`)
	validEyeColors := "amb blu brn gry grn hzl oth"

	passports := make([]map[string]string, len(values))
	for i, v := range values {
		fields := strings.Fields(v)
		p := make(map[string]string)
		for _, f := range fields {
			splitField := strings.Split(f, ":")
			p[splitField[0]] = splitField[1]
		}
		passports[i] = p
	}

	validCount := 0
	for _, p := range passports {
		if !all(p, requiredFieldFunctions) {
			continue
		}
		birthYear, _ := strconv.Atoi(p["byr"])
		if birthYear > 2002 || birthYear < 1920 {
			continue
		}
		issueYear, _ := strconv.Atoi(p["iyr"])
		if issueYear > 2020 || issueYear < 2010 {
			continue
		}
		expirationYear, _ := strconv.Atoi(p["eyr"])
		if expirationYear > 2030 || expirationYear < 2020 {
			continue
		}
		heightMatch := heightRegex.FindStringSubmatch(p["hgt"])
		if heightMatch == nil {
			continue
		}
		if heightMatch[2] != "in" && heightMatch[2] != "cm" {
			continue
		}
		if heightMatch[2] == "cm" {
			height, _ := strconv.Atoi(heightMatch[1])
			if height < 150 || height > 193 {
				continue
			}
		}
		if heightMatch[2] == "in" {
			height, _ := strconv.Atoi(heightMatch[1])
			if height < 59 || height > 76 {
				continue
			}
		}
		if !hairColorRegex.MatchString(p["hcl"]) {
			continue
		}
		if !strings.Contains(validEyeColors, p["ecl"]) {
			continue
		}
		if !passportIDRegex.MatchString(p["pid"]) {
			continue
		}
		validCount++
		fmt.Println(p)
	}
	return validCount
}
