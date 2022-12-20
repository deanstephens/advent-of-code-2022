package day18

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() int {
	coords := read(false)
	fmt.Println(coords)

	totalSides := len(coords) * 6

	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			if coords[i].IsAdjacent(coords[j]) {
				totalSides -= 2
			}
		}
	}
	return totalSides
}

func Part2() int {
	coords := read(false)
	fmt.Println(coords)

	maxCoord := utils.Coordinate3d{
		X: 0,
		Y: 0,
		Z: 0,
	}
	minCoord := utils.Coordinate3d{
		X: -1,
		Y: -1,
		Z: -1,
	}
	for _, coord := range coords {
		if coord.X > maxCoord.X {
			maxCoord.X = coord.X
		}
		if coord.Y > maxCoord.Y {
			maxCoord.Y = coord.Y
		}
		if coord.Z > maxCoord.Z {
			maxCoord.Z = coord.Z
		}
	}
	totalSides := len(coords) * 6

	maxCoord.X++
	maxCoord.Y++
	maxCoord.Z++

	visitedPositions := map[string]bool{}

	visitedCount := navigate(minCoord, visitedPositions, minCoord, maxCoord, coords)

	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			if coords[i].IsAdjacent(coords[j]) {
				totalSides -= 2
			}
		}
	}
	return visitedCount
}

func navigate(coord utils.Coordinate3d, visitedPositions map[string]bool, minCoord utils.Coordinate3d, maxCoord utils.Coordinate3d, cubes []utils.Coordinate3d) int {
	totalVisitedSides := 0
	if coord.X < minCoord.X || coord.Y < minCoord.Y || coord.Z < minCoord.Z ||
		coord.X > maxCoord.X || coord.Y > maxCoord.Y || coord.Z > maxCoord.Z {
		return 0
	}

	if _, ok := visitedPositions[fmt.Sprintf("%d,%d,%d", coord.X, coord.Y, coord.Z)]; ok {
		return 0
	}

	for _, cube := range cubes {
		if cube.Equals(coord) {
			return 1
		}
	}

	visitedPositions[fmt.Sprintf("%d,%d,%d", coord.X, coord.Y, coord.Z)] = true

	totalVisitedSides += navigate(coord.Add(utils.Coordinate3d{X: 0, Y: 0, Z: 1}), visitedPositions, minCoord, maxCoord, cubes)
	totalVisitedSides += navigate(coord.Add(utils.Coordinate3d{X: 0, Y: 0, Z: -1}), visitedPositions, minCoord, maxCoord, cubes)
	totalVisitedSides += navigate(coord.Add(utils.Coordinate3d{X: 0, Y: 1, Z: 0}), visitedPositions, minCoord, maxCoord, cubes)
	totalVisitedSides += navigate(coord.Add(utils.Coordinate3d{X: 0, Y: -1, Z: 0}), visitedPositions, minCoord, maxCoord, cubes)
	totalVisitedSides += navigate(coord.Add(utils.Coordinate3d{X: 1, Y: 0, Z: 0}), visitedPositions, minCoord, maxCoord, cubes)
	totalVisitedSides += navigate(coord.Add(utils.Coordinate3d{X: -1, Y: 0, Z: 0}), visitedPositions, minCoord, maxCoord, cubes)

	return totalVisitedSides
}

func read(testInput bool) []utils.Coordinate3d {
	fileName := "input.txt"
	if testInput {
		fileName = "test.txt"
	}
	file, err := os.Open("day18/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	coords := []utils.Coordinate3d{}

	for scanner.Scan() {
		text := scanner.Text()

		splitText := strings.Split(text, ",")
		x, _ := strconv.Atoi(splitText[0])
		y, _ := strconv.Atoi(splitText[1])
		z, _ := strconv.Atoi(splitText[2])

		coords = append(coords, utils.Coordinate3d{X: x, Y: y, Z: z})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return coords
}
