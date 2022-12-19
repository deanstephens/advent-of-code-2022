package day17

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type displayable interface {
	display() string
}

type rock struct{}

func (r rock) display() string {
	return "#"
}

type fallingRock struct {
	coords []utils.Coordinate
	height int
}

func (r fallingRock) display() string {
	return "@"
}

type flatState struct {
	fallingRockIndex int
	jetPatternIndex  int
	height           int
}

type chamber struct {
	fallingRock      fallingRock
	fallingRockIndex int
	grid             [][]displayable
	highestRock      int
	fallingRocks     []fallingRock
	jetPatterns      []direction
	jetPatternIndex  int
	flats            map[string]flatState
	lastFlat         int
}

type direction string

const leftDirection direction = "<"
const rightDirection direction = ">"

func (c chamber) getDisplayAtCoord(coord utils.Coordinate) string {
	if c.fallingRock.coords != nil {
		for _, frc := range c.fallingRock.coords {
			if frc.X == coord.X && frc.Y == coord.Y {
				return "@"
			}
		}
	}

	gridEle := c.grid[coord.Y][coord.X]
	if gridEle != nil {
		return gridEle.display()
	}

	return "."
}

func (c chamber) display(rows int) {
	fmt.Print("\033[H\033[2J")

	bottomIndex := 0
	if len(c.grid)-rows > 0 && rows != -1 {
		bottomIndex = len(c.grid) - rows
	}

	for i := len(c.grid) - 1; i >= bottomIndex; i-- {
		fmt.Print("|")
		for j, _ := range c.grid[i] {
			fmt.Print(c.getDisplayAtCoord(utils.Coordinate{X: j, Y: i}))
		}
		fmt.Print("|\n")
		if i == 0 {
			fmt.Print("+-------+\n")
		}
	}
}

func (c *chamber) extendGrid(numRows int) {
	for i := 0; i < numRows; i++ {
		c.grid = append(c.grid, make([]displayable, 7))
	}
}

func (c *chamber) init() {
	c.grid = make([][]displayable, 4)

	for row := 0; row < 4; row++ {
		c.grid[row] = make([]displayable, 7)
	}

	c.highestRock = 0

	c.fallingRocks = []fallingRock{
		{
			coords: []utils.Coordinate{{
				X: 0,
				Y: 0,
			}, {
				X: 1,
				Y: 0,
			}, {
				X: 2,
				Y: 0,
			}, {
				X: 3,
				Y: 0,
			}},
			height: 1,
		},
		{
			coords: []utils.Coordinate{{
				X: 0,
				Y: 1,
			}, {
				X: 1,
				Y: 2,
			}, {
				X: 1,
				Y: 1,
			}, {
				X: 1,
				Y: 0,
			}, {
				X: 2,
				Y: 1,
			}},
			height: 3,
		},
		{
			coords: []utils.Coordinate{{
				X: 0,
				Y: 0,
			}, {
				X: 1,
				Y: 0,
			}, {
				X: 2,
				Y: 0,
			}, {
				X: 2,
				Y: 1,
			}, {
				X: 2,
				Y: 2,
			}},
			height: 3,
		},
		{
			coords: []utils.Coordinate{{
				X: 0,
				Y: 0,
			}, {
				X: 0,
				Y: 1,
			}, {
				X: 0,
				Y: 2,
			}, {
				X: 0,
				Y: 3,
			}},
			height: 4,
		},
		{
			coords: []utils.Coordinate{{
				X: 0,
				Y: 0,
			}, {
				X: 1,
				Y: 0,
			}, {
				X: 0,
				Y: 1,
			}, {
				X: 1,
				Y: 1,
			}},
			height: 2,
		},
	}
}

func Part1() int {
	c := chamber{}

	c.jetPatterns = read(false)

	c.init()

	for i := 0; i < 2022; i++ {
		c.dropRock(false, 20)
	}

	return 0
}

func Part2() int {
	c := chamber{}

	c.jetPatterns = read(true)

	c.init()

	for i := 0; i < 1000000000000; i++ {
		if i%100000 == 0 {
			fmt.Println(i)
		}
		c.dropRock(true, 30)
		if c.isFlat() {
			fmt.Println("Is flat", i)
		}
	}

	return 0
}

func (c *chamber) jetBlastRock() {
	if c.jetPatternIndex >= len(c.jetPatterns) {
		c.jetPatternIndex = 0
	}

	nextMove := c.jetPatterns[c.jetPatternIndex]
	c.jetPatternIndex++

	switch nextMove {
	case leftDirection:
		c.moveFallingRock(utils.Coordinate{X: -1, Y: 0})
		break
	case rightDirection:
		c.moveFallingRock(utils.Coordinate{X: 1, Y: 0})
		break
	}
}

func (c *chamber) moveFallingRock(offset utils.Coordinate) bool {
	movedCoords := []utils.Coordinate{}
	for _, coord := range c.fallingRock.coords {
		newCoord := coord.Add(offset)
		if newCoord.X < 0 || newCoord.X >= 7 || newCoord.Y < 0 || c.grid[newCoord.Y][newCoord.X] != nil {
			return false
		}
		movedCoords = append(movedCoords, newCoord)
	}

	c.fallingRock.coords = movedCoords
	return true
}

func (c chamber) isFlat() bool {
	return c.grid[c.highestRock-1][0] != nil && c.grid[c.highestRock-1][1] != nil && c.grid[c.highestRock-1][2] != nil &&
		c.grid[c.highestRock-1][3] != nil && c.grid[c.highestRock-1][4] != nil && c.grid[c.highestRock-1][5] != nil &&
		c.grid[c.highestRock-1][6] != nil
}

func (c *chamber) dropRock(display bool, displayRows int) {
	if c.fallingRockIndex >= len(c.fallingRocks) {
		c.fallingRockIndex = 0
	}
	c.fallingRock = c.fallingRocks[c.fallingRockIndex]

	c.fallingRockIndex++

	requiredHeight := (c.fallingRock.height + 3 + c.highestRock) - len(c.grid)
	if requiredHeight > 0 {
		c.extendGrid(requiredHeight)
	}

	startingOffset := utils.Coordinate{X: 2, Y: c.highestRock + 3}
	startingCoords := []utils.Coordinate{}
	for _, coord := range c.fallingRock.coords {
		startingCoords = append(startingCoords, utils.Coordinate{X: coord.X + startingOffset.X, Y: coord.Y + startingOffset.Y})
	}

	c.fallingRock.coords = startingCoords

	if display {
		c.display(displayRows)
		time.Sleep(100 * time.Millisecond)
	}

	for {
		c.jetBlastRock()

		if display {
			c.display(displayRows)
			time.Sleep(100 * time.Millisecond)
		}

		isFalling := c.moveFallingRock(utils.Coordinate{X: 0, Y: -1})

		if display {
			c.display(displayRows)
			time.Sleep(100 * time.Millisecond)
		}

		if !isFalling {
			for _, coord := range c.fallingRock.coords {
				c.grid[coord.Y][coord.X] = rock{}
				if coord.Y+1 > c.highestRock {
					c.highestRock = coord.Y + 1
				}
			}
			break
		}
	}
}

func read(testInput bool) []direction {
	fileName := "input.txt"
	if testInput {
		fileName = "test.txt"
	}
	file, err := os.Open("day17/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	directions := []direction{}

	for scanner.Scan() {
		text := scanner.Text()

		splitText := strings.Split(text, "")

		for _, s := range splitText {
			newDirection := leftDirection
			if s == ">" {
				newDirection = rightDirection
			}
			directions = append(directions, newDirection)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return directions
}
