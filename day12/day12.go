package day12

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Part1() int {
	grid, stepsGrid, start, end := read(false)

	goalPath := navigate(start, end, grid, stepsGrid, 0)

	printPath(grid, goalPath)
	return len(goalPath) - 1
}

func Part2() int {
	grid, stepsGrid, start, end := read(false)

	bestGoalPath := []utils.Coordinate{}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 1 {
				start = utils.Coordinate{X: col, Y: row}
				goalPath := navigate(start, end, grid, stepsGrid, 0)
				if len(goalPath) != 0 && (len(goalPath) < len(bestGoalPath) || len(bestGoalPath) == 0) {
					bestGoalPath = goalPath
				}
			}
		}
	}

	return len(bestGoalPath) - 1
}

func printPath(grid [][]int, visited []utils.Coordinate) {
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			visit := false
			for c, coord := range visited {
				if c == len(visited)-1 {
					continue
				}
				if coord.Y == row && coord.X == column {
					visit = true
					if coord.Y > visited[c+1].Y {
						fmt.Print("^")
					}
					if coord.Y < visited[c+1].Y {
						fmt.Print("v")
					}
					if coord.X > visited[c+1].X {
						fmt.Print("<")
					}
					if coord.X < visited[c+1].X {
						fmt.Print(">")
					}
				}
			}
			if !visit {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

	fmt.Print("\n\n\n")
}

func navigationScoreToNode(coord utils.Coordinate, prevCoord utils.Coordinate, grid [][]int, minStepsGrid [][]*int, steps int) (res *int) {
	if coord.Y < 0 || coord.Y >= len(grid) || coord.X < 0 || coord.X >= len(grid[coord.Y]) {
		return res
	}

	prevHeight := grid[prevCoord.Y][prevCoord.X]
	elevationGain := grid[coord.Y][coord.X] - prevHeight

	if elevationGain > 1 {
		return res
	}

	if minStepsGrid[coord.Y][coord.X] != nil && *minStepsGrid[coord.Y][coord.X] <= steps {
		return res
	}

	if steps >= minStepsToFinish {
		return res
	}

	res = &elevationGain

	return res
}

func lowestPathByLength(paths [][]utils.Coordinate) []utils.Coordinate {
	lowestLengthPath := []utils.Coordinate{}

	for _, path := range paths {
		if len(path) > 0 && (len(path) < len(lowestLengthPath) || len(lowestLengthPath) == 0) {
			lowestLengthPath = path
		}
	}

	return lowestLengthPath
}

var minStepsToFinish = 100000000000

func navigate(current utils.Coordinate, goal utils.Coordinate, grid [][]int, minStepsGrid [][]*int, steps int) []utils.Coordinate {
	minStepsGrid[current.Y][current.X] = &steps
	nextSteps := steps + 1
	if current.X == goal.X && current.Y == goal.Y {
		if steps < minStepsToFinish {
			minStepsToFinish = steps
		}
		return []utils.Coordinate{{X: current.X, Y: current.Y}}
	}
	paths := [][]utils.Coordinate{}

	upCoord := utils.Coordinate{X: current.X, Y: current.Y - 1}
	if navigationScoreToNode(upCoord, current, grid, minStepsGrid, nextSteps) != nil {
		pathToEnd := navigate(upCoord, goal, grid, minStepsGrid, nextSteps)
		if len(pathToEnd) != 0 {
			paths = append(paths, pathToEnd)
		}
	}
	downCoord := utils.Coordinate{X: current.X, Y: current.Y + 1}
	if navigationScoreToNode(downCoord, current, grid, minStepsGrid, nextSteps) != nil {
		pathToEnd := navigate(downCoord, goal, grid, minStepsGrid, nextSteps)
		if len(pathToEnd) != 0 {
			paths = append(paths, pathToEnd)
		}
	}
	leftCoord := utils.Coordinate{X: current.X - 1, Y: current.Y}
	if navigationScoreToNode(leftCoord, current, grid, minStepsGrid, nextSteps) != nil {
		pathToEnd := navigate(leftCoord, goal, grid, minStepsGrid, nextSteps)
		if len(pathToEnd) != 0 {
			paths = append(paths, pathToEnd)
		}
	}
	rightCoord := utils.Coordinate{X: current.X + 1, Y: current.Y}
	if navigationScoreToNode(rightCoord, current, grid, minStepsGrid, nextSteps) != nil {
		pathToEnd := navigate(rightCoord, goal, grid, minStepsGrid, nextSteps)
		if len(pathToEnd) != 0 {
			paths = append(paths, pathToEnd)
		}
	}

	pathToEnd := lowestPathByLength(paths)

	if len(pathToEnd) > 0 {
		return append([]utils.Coordinate{current}, pathToEnd...)
	}

	return []utils.Coordinate{}
}

func read(testInput bool) (grid [][]int, stepsGrid [][]*int, start utils.Coordinate, end utils.Coordinate) {
	fileName := "input.txt"
	if testInput {
		fileName = "test.txt"
	}
	file, err := os.Open("day12/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		rowChars := strings.Split(text, "")
		rowInts := []int{}

		var stepsSlice = make([]*int, len(rowChars))

		for i, rowChar := range rowChars {
			if rowChar == "S" {
				rowInts = append(rowInts, 1)
				start = utils.Coordinate{
					X: i,
					Y: len(grid),
				}
				continue
			}
			if rowChar == "E" {
				rowInts = append(rowInts, 26)
				end = utils.Coordinate{
					X: i,
					Y: len(grid),
				}
				continue
			}

			rowInts = append(rowInts, int(rowChar[0])-96)
		}

		grid = append(grid, rowInts)
		stepsGrid = append(stepsGrid, stepsSlice)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid, stepsGrid, start, end
}
