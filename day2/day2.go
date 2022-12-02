package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type round struct {
	them  string
	you   string
	score int
}

func Day2Part1() int {
	rounds := read()
	fmt.Println(rounds)

	totalScore := 0

	for _, round := range rounds {
		totalScore += round.score
	}

	return totalScore
}

func Day2Part2() int {
	rounds := read()
	fmt.Println(rounds)

	totalScore := 0

	for _, round := range rounds {
		newThrow := getRequiredThrow(round.them, round.you)
		totalScore += getScore(round.them, newThrow)
	}

	return totalScore
}

func read() []round {
	rounds := []round{}

	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	for scanner.Scan() {
		text := scanner.Text()

		s := strings.Split(text, " ")
		rounds = append(rounds, round{them: s[0], you: s[1], score: getScore(s[0], s[1])})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rounds
}

const rock = "A"
const youRock = "X"

const paper = "B"
const youPaper = "Y"

const scissors = "C"
const youScissors = "Z"

func getRequiredThrow(them string, outcome string) string {
	if outcome == "X" { //lose
		if them == rock {
			return youScissors
		}
		if them == paper {
			return youRock
		}
		if them == scissors {
			return youPaper
		}
	}

	if outcome == "Y" { //draw
		if them == rock {
			return youRock
		}
		if them == paper {
			return youPaper
		}
		if them == scissors {
			return youScissors
		}
	}

	if outcome == "Z" { //win
		if them == rock {
			return youPaper
		}
		if them == paper {
			return youScissors
		}
		if them == scissors {
			return youRock
		}
	}

	return ""
}

func getScore(them string, you string) int {
	if you == youRock {
		if them == rock {
			return 1 + 3
		}
		if them == paper {
			return 1 + 0
		}
		if them == scissors {
			return 1 + 6
		}
	}

	if you == youPaper {
		if them == rock {
			return 2 + 6
		}
		if them == paper {
			return 2 + 3
		}
		if them == scissors {
			return 2 + 0
		}
	}

	if you == youScissors {
		if them == rock {
			return 3 + 0
		}
		if them == paper {
			return 3 + 6
		}
		if them == scissors {
			return 3 + 3
		}
	}

	return 0
}
