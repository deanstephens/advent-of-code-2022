package day14

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type grid struct {
	obstructionList [][]*obstruction
	minX            int
	maxX            int
	minY            int
	maxY            int
	offsetX         int
}

func (g *grid) initGridSize() {
	g.obstructionList = make([][]*obstruction, g.maxY+4)

	for i := 0; i < g.maxY+4; i++ {
		g.obstructionList[i] = make([]*obstruction, g.maxX+1)
	}
}

func (g *grid) resizeGridIfNeeded(a utils.Coordinate) {
	if a.X < g.minX {
		g.minX = a.X
	}
	if a.X > g.maxX {
		g.maxX = a.X
	}
	if a.X+g.offsetX < 0 {
		for i := 0; i < len(g.obstructionList); i++ {
			g.obstructionList[i] = append(make([]*obstruction, 1), g.obstructionList[i]...)
		}
		g.offsetX++
	}
	if a.X+g.offsetX > len(g.obstructionList[0])-1 {
		for i := 0; i < len(g.obstructionList); i++ {
			var obstructionToInsert *obstruction
			if i == len(g.obstructionList)-2 {
				obstructionToInsert = &obstruction{display: "#"}
			}
			g.obstructionList[i] = append(g.obstructionList[i], obstructionToInsert)
		}
	}
}

func (g grid) getObstructionAtCoord(y int, x int) *obstruction {
	return g.obstructionList[y][x+g.offsetX]
}

func (g grid) setObstructionAtCoord(y int, x int, o *obstruction) {
	g.obstructionList[y][x+g.offsetX] = o
}

func (g *grid) display() {
	fmt.Print("\033[H\033[2J")
	for y := g.minY; y <= g.maxY+3; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			obstruction := g.getObstructionAtCoord(y, x)
			if obstruction != nil {
				fmt.Print(obstruction.display)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

type obstruction struct {
	display string
}

func Part1() int {
	theGrid := read(true)
	theGrid.display()

	sandPiecesDropped := 0
	for {
		didSettle := theGrid.dropSand()
		if !didSettle {
			break
		}
		sandPiecesDropped++
	}

	return sandPiecesDropped
}

func Part2() int {
	theGrid := read(false)
	//theGrid.display()

	for i := 0; i <= theGrid.maxX; i++ {
		theGrid.setObstructionAtCoord(theGrid.maxY+2, i, &obstruction{display: "#"})
	}

	sandPiecesDropped := 0
	for {
		didSettle := theGrid.dropSand()
		if !didSettle {
			break
		}
		sandPiecesDropped++
	}
	theGrid.display()

	return sandPiecesDropped
}

func (g *grid) dropSand() bool {
	sandCoord := utils.Coordinate{X: 500, Y: 0}
	if g.obstructionList[sandCoord.Y][sandCoord.X] != nil {
		return false
	}
	fallingSand := obstruction{display: "+"}
	g.setObstructionAtCoord(sandCoord.Y, sandCoord.X, &fallingSand)

	for {
		if sandCoord.Y+1 > len(g.obstructionList)-1 {
			return false
		}
		//g.display()
		if g.getObstructionAtCoord(sandCoord.Y+1, sandCoord.X) == nil {
			g.setObstructionAtCoord(sandCoord.Y, sandCoord.X, nil)
			g.setObstructionAtCoord(sandCoord.Y+1, sandCoord.X, &fallingSand)
			sandCoord.Y = sandCoord.Y + 1
		} else {
			g.resizeGridIfNeeded(utils.Coordinate{X: sandCoord.X - 1, Y: sandCoord.Y + 1})
			g.resizeGridIfNeeded(utils.Coordinate{X: sandCoord.X + 1, Y: sandCoord.Y + 1})
			if g.getObstructionAtCoord(sandCoord.Y+1, sandCoord.X-1) == nil {
				g.setObstructionAtCoord(sandCoord.Y, sandCoord.X, nil)
				g.setObstructionAtCoord(sandCoord.Y+1, sandCoord.X-1, &fallingSand)
				sandCoord.Y = sandCoord.Y + 1
				sandCoord.X = sandCoord.X - 1
			} else if g.getObstructionAtCoord(sandCoord.Y+1, sandCoord.X+1) == nil {
				g.setObstructionAtCoord(sandCoord.Y, sandCoord.X, nil)
				g.setObstructionAtCoord(sandCoord.Y+1, sandCoord.X+1, &fallingSand)
				sandCoord.Y = sandCoord.Y + 1
				sandCoord.X = sandCoord.X + 1
			} else {
				g.setObstructionAtCoord(sandCoord.Y, sandCoord.X, &obstruction{display: "o"})
				return true
			}
		}
	}
}

func read(testInput bool) grid {
	fileName := "input.txt"
	if testInput {
		fileName = "test.txt"
	}
	file, err := os.Open("day14/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	theGrid := grid{}

	rockCoords := []utils.Coordinate{}

	theGrid.maxX = 500
	theGrid.minX = 1000000
	theGrid.maxY = 0
	theGrid.minY = 0

	for scanner.Scan() {
		text := scanner.Text()

		points := strings.Split(text, " -> ")

		pointsAsCoords := []utils.Coordinate{}

		for _, p := range points {
			coords := strings.Split(p, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])

			if x > theGrid.maxX {
				theGrid.maxX = x
			}
			if y > theGrid.maxY {
				theGrid.maxY = y
			}
			if x < theGrid.minX {
				theGrid.minX = x
			}
			if y < theGrid.minY {
				theGrid.minY = y
			}

			pointsAsCoords = append(pointsAsCoords, utils.Coordinate{
				X: x,
				Y: y,
			})
		}

		for i := 0; i < len(pointsAsCoords)-1; i++ {
			p1 := pointsAsCoords[i]
			p2 := pointsAsCoords[i+1]
			xMag, xDir := utils.GetMagnitudeAndDirection2d(p1.X, p2.X)
			yMag, yDir := utils.GetMagnitudeAndDirection2d(p1.Y, p2.Y)
			for x := 0; x < xMag; x++ {
				rockCoords = append(rockCoords, utils.Coordinate{X: pointsAsCoords[i].X + (x * xDir), Y: pointsAsCoords[i].Y})
			}
			for y := 0; y < yMag; y++ {
				rockCoords = append(rockCoords, utils.Coordinate{X: pointsAsCoords[i].X, Y: pointsAsCoords[i].Y + (y * yDir)})
			}
		}
		rockCoords = append(rockCoords, pointsAsCoords[len(pointsAsCoords)-1])
	}

	theGrid.initGridSize()

	for _, coord := range rockCoords {
		theGrid.obstructionList[coord.Y][coord.X] = &obstruction{display: "#"}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return theGrid
}
