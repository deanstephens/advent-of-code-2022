package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type operation struct {
	operator string
	arg      int
	cycles   int
}

type coordinate struct {
	x int
	y int
}

type knot struct {
	currentPosition  coordinate
	visitedPositions []coordinate
}

func Part1() int {
	operations := read()

	x := 1
	clockCycle := 1

	signalStrength := 0

	for _, op := range operations {
		for i := 1; i <= op.cycles; i++ {
			if clockCycle == 20 || (clockCycle+20)%40 == 0 {
				additionalSignalStrength := clockCycle * x
				signalStrength += additionalSignalStrength
			}
			clockCycle++
		}
		x += op.arg
	}

	return signalStrength
}

func Part2() int {
	operations := read()

	x := 1
	clockCycle := 0

	for _, op := range operations {
		for i := 1; i <= op.cycles; i++ {
			if clockCycle%40 == 0 && clockCycle != 0 {
				fmt.Print("\n")
			}
			if clockCycle%40 <= x+1 && clockCycle%40 >= x-1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			clockCycle++
		}
		x += op.arg
	}

	return 0
}

func read() []operation {
	file, err := os.Open("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	operations := []operation{}

	for scanner.Scan() {
		text := scanner.Text()
		c := strings.Split(text, " ")

		argument := 0
		clockCycles := 1

		if c[0] == "addx" {
			argument, _ = strconv.Atoi(c[1])
			clockCycles = 2
		}

		operations = append(operations, operation{
			operator: c[0],
			arg:      argument,
			cycles:   clockCycles,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return operations
}
