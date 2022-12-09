package day9

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	distance  int
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
	commands := read()
	head := knot{
		currentPosition: coordinate{
			x: 0,
			y: 0,
		},
		visitedPositions: []coordinate{{
			x: 0,
			y: 0,
		}},
	}
	tail := knot{
		currentPosition: coordinate{
			x: 0,
			y: 0,
		},
		visitedPositions: []coordinate{{
			x: 0,
			y: 0,
		}},
	}

	for _, com := range commands {
		followCommand(com, &head, []*knot{&tail})
	}

	visitedTailCoords := map[coordinate]bool{}

	for _, coord := range tail.visitedPositions {
		visitedTailCoords[coord] = true
	}

	return len(visitedTailCoords)
}

func Part2() int {
	commands := read()

	head := knot{
		currentPosition: coordinate{
			x: 0,
			y: 0,
		},
		visitedPositions: []coordinate{{
			x: 0,
			y: 0,
		}},
	}

	knots := []*knot{}
	for i := 0; i < 9; i++ {
		knots = append(knots, &knot{
			currentPosition: coordinate{
				x: 0,
				y: 0,
			},
			visitedPositions: []coordinate{{
				x: 0,
				y: 0,
			}},
		})
	}

	for _, com := range commands {
		followCommand(com, &head, knots)
		printCurrentState(head, knots)
	}

	visitedTailCoords := map[coordinate]bool{}

	for _, coord := range knots[8].visitedPositions {
		visitedTailCoords[coord] = true
	}

	displayGrid(visitedTailCoords)
	return len(visitedTailCoords)
}

func displayGrid(coordMap map[coordinate]bool) {
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for coord, _ := range coordMap {
		if coord.x < minX {
			minX = coord.x
		}
		if coord.x > maxX {
			maxX = coord.x
		}

		if coord.y < minY {
			minY = coord.y
		}
		if coord.y > maxY {
			maxY = coord.y
		}
	}

	for row := minY; row < maxY; row++ {
		for col := minX; col < maxX; col++ {
			_, ok := coordMap[coordinate{x: col, y: row}]
			if ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func followCommand(com command, head *knot, tails []*knot) {
	for i := 0; i < com.distance; i++ {
		performHeadStep(com.direction, head)
		performTailStep(head, tails[0])

		for j := 1; j < len(tails); j++ {
			performTailStep(tails[j-1], tails[j])
		}
	}
}

func printCurrentState(head knot, tails []*knot) {
	for row := 15; row >= -5; row-- {
		for col := -11; col <= 14; col++ {
			printed := false
			if head.currentPosition.x == col && head.currentPosition.y == row {
				fmt.Print("H")
				continue
			}
			for i, tail := range tails {
				if tail.currentPosition.x == col && tail.currentPosition.y == row {
					fmt.Print(i + 1)
					printed = true
					break
				}
			}
			if !printed {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("")
}

func performHeadStep(direction string, head *knot) {
	switch direction {
	case "U":
		head.currentPosition = coordinate{
			x: head.currentPosition.x,
			y: head.currentPosition.y + 1,
		}
		break
	case "D":
		head.currentPosition = coordinate{
			x: head.currentPosition.x,
			y: head.currentPosition.y - 1,
		}
		break
	case "L":
		head.currentPosition = coordinate{
			x: head.currentPosition.x - 1,
			y: head.currentPosition.y,
		}
		break
	case "R":
		head.currentPosition = coordinate{
			x: head.currentPosition.x + 1,
			y: head.currentPosition.y,
		}
		break
	}
	head.visitedPositions = append(head.visitedPositions, head.currentPosition)
}

func performTailStep(head *knot, tail *knot) {
	diffY := head.currentPosition.y - tail.currentPosition.y
	diffX := head.currentPosition.x - tail.currentPosition.x

	directionX := 0
	if diffX != 0 {
		directionX = diffX / int(math.Abs(float64(diffX)))
	}
	directionY := 0
	if diffY != 0 {
		directionY = diffY / int(math.Abs(float64(diffY)))
	}

	moved := true
	if diffY > 1 || diffX > 1 || diffY < -1 || diffX < -1 {
		tail.currentPosition = coordinate{
			x: tail.currentPosition.x + directionX,
			y: tail.currentPosition.y + directionY,
		}

		moved = true
	}

	if moved {
		tail.visitedPositions = append(tail.visitedPositions, tail.currentPosition)
	}
}

func read() []command {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	commands := []command{}

	for scanner.Scan() {
		text := scanner.Text()
		c := strings.Split(text, " ")
		distance, _ := strconv.Atoi(c[1])

		commands = append(commands, command{
			direction: c[0],
			distance:  distance,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return commands
}
