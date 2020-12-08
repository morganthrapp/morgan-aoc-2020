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

// Instruction represents one instruction
type Instruction struct {
	instructionType string
	runCount        int
	argument        int
	run             func(Context) Context
}

// Context is the global program context
type Context struct {
	instructions       []Instruction
	instructionPointer int
	accumulator        int
}

func main() {
	rawData, err := ioutil.ReadFile("input")
	check(err)
	data := strings.Split(string(rawData), "\n")

	context := Context{
		instructionPointer: 0,
		accumulator:        0,
		instructions:       make([]Instruction, len(data)),
	}
	instructionRegex, _ := regexp.Compile(`(\w{3}) ([+-]\d+)`)
	for i, line := range data {
		instructionMatch := instructionRegex.FindStringSubmatch(line)
		instructionArgument, _ := strconv.Atoi(instructionMatch[2])
		switch instructionMatch[1] {
		case "nop":
			context.instructions[i] = buildNoop(instructionArgument)
		case "jmp":
			context.instructions[i] = buildJump(instructionArgument)
		case "acc":
			context.instructions[i] = buildAcc(instructionArgument)

		}
	}

	fmt.Println(problemOne(context))
	fmt.Println(problemTwo(context))
}

func buildNoop(argument int) Instruction {
	return Instruction{
		runCount:        0,
		instructionType: "nop",
		argument:        argument,
		run: func(c Context) Context {
			return Context{
				instructions:       c.instructions,
				instructionPointer: c.instructionPointer + 1,
				accumulator:        c.accumulator,
			}
		},
	}
}

func buildJump(argument int) Instruction {
	return Instruction{
		runCount:        0,
		instructionType: "jmp",
		argument:        argument,
		run: func(c Context) Context {
			return Context{
				instructions:       c.instructions,
				instructionPointer: c.instructionPointer + argument,
				accumulator:        c.accumulator,
			}
		},
	}
}

func buildAcc(argument int) Instruction {
	return Instruction{
		runCount:        0,
		instructionType: "acc",
		argument:        argument,
		run: func(c Context) Context {
			return Context{
				instructions:       c.instructions,
				instructionPointer: c.instructionPointer + 1,
				accumulator:        c.accumulator + argument,
			}
		},
	}
}

func (c Context) next() Instruction {
	return c.instructions[c.instructionPointer]
}

func problemOne(context Context) int {
	instruction := context.next()
	for instruction.runCount < 2 {
		context = instruction.run(context)
		instruction = context.next()
		instruction.runCount++
		context.instructions[context.instructionPointer] = instruction
	}
	return context.accumulator
}

func problemTwo(context Context) int {
	brokenContext := context
	instruction := brokenContext.next()
	for instruction.runCount < 2{
		newContext := instruction.run(brokenContext)
		if newContext.instructionPointer >= len(newContext.instructions) {
			return newContext.accumulator
		}
		instruction.runCount++
		brokenContext.instructions[brokenContext.instructionPointer] = instruction
		instruction = newContext.next()
		if instruction.runCount > 0 {
			brokenInstructionAddress := brokenContext.instructionPointer
			brokenInstruction := brokenContext.next()
			if brokenInstruction.instructionType == "jmp" {
				context.instructions[brokenInstructionAddress] = buildNoop(brokenInstruction.argument)
			} else {
				context.instructions[brokenInstructionAddress] = buildJump(brokenInstruction.argument)
			}
		}
		brokenContext = newContext
	}
	return problemTwo(context)
}
