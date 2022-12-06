package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack []string

func Part1() string {
	stacks, instructions := read()
	fmt.Println(stacks, instructions)

	stacksTop := ""
	for _, instruction := range instructions {
		stacks = followInstruction(stacks, instruction, true)
	}

	for _, stack := range stacks {
		stacksTop += stack[len(stack)-1]
	}

	fmt.Println(stacksTop)

	return stacksTop
}

func Part2() string {
	stacks, instructions := read()
	fmt.Println(stacks, instructions)

	stacksTop := ""
	for _, instruction := range instructions {
		stacks = followInstruction(stacks, instruction, false)
	}

	for _, stack := range stacks {
		stacksTop += stack[len(stack)-1]
	}

	fmt.Println(stacksTop)

	return stacksTop
}

func followInstruction(stacks []stack, instruction instruction, doReverse bool) []stack {
	newStack := stacks[:]
	from := stacks[instruction.from]
	for i := 0; i < instruction.count; i++ {
	}
	forMoving := from[len(from)-instruction.count:]
	if doReverse {
		reverse(forMoving)
	}

	newStack[instruction.from] = from[:len(from)-instruction.count]
	newStack[instruction.to] = append(newStack[instruction.to], forMoving...)

	return newStack
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func getStacksFromStrings(stackStrings []string) []stack {
	stacks := []stack{}

	for i := len(stackStrings) - 1; i >= 0; i-- {
		for j := 0; j < (len(stackStrings[i])+1)/4; j++ {
			char := stackStrings[i][(j*4)+1 : (j*4)+2]
			if char != " " {
				if len(stacks) < (j + 1) {
					stacks = append(stacks, stack{})
				}
				stacks[j] = append(stacks[j], char)
			}
		}
	}

	return stacks
}

type instruction struct {
	count int
	from  int
	to    int
}

func read() ([]stack, []instruction) {
	var stacks []stack

	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var stackStrings = []string{}

	readingStacks := true

	instructions := []instruction{}

	// optionally, resize scanner's capacity for lines over 64K, see next example https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		}

		if readingStacks {
			if text[1:2] == "1" {
				readingStacks = false
				stacks = getStacksFromStrings(stackStrings)
				continue
			}
			stackStrings = append(stackStrings, text)
			continue
		}

		instructionString := strings.Replace(text, "move ", "", 1)
		instructionString = strings.Replace(instructionString, "from ", "", 1)
		instructionString = strings.Replace(instructionString, "to ", "", 1)

		instructionArr := strings.Split(instructionString, " ")
		iCount, _ := strconv.Atoi(instructionArr[0])
		iFrom, _ := strconv.Atoi(instructionArr[1])
		iTo, _ := strconv.Atoi(instructionArr[2])

		instructions = append(instructions, instruction{
			count: iCount,
			from:  iFrom - 1,
			to:    iTo - 1,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return stacks, instructions
}
