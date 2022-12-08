package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type tree struct {
	height    int
	isVisible bool
}

func Part1() int {
	trees := read()
	visibleTreeCount := updateVisibleTrees(trees)

	return visibleTreeCount
}

func Part2() int {
	trees := read()

	maxScenicScore := 0
	for row := 0; row < len(trees); row++ {
		for col := 0; col < len(trees[row]); col++ {
			scenicScore := getScenicScore(trees, row, col)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	return maxScenicScore
}

func printVisibility(trees [][]*tree) {
	for row := 0; row < len(trees); row++ {
		for col := 0; col < len(trees[row]); col++ {
			if trees[row][col].isVisible {
				fmt.Print("V")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println("")
	}
}

func isTreeVisible(trees [][]*tree, row int, col int) bool {
	if col == 0 || row == 0 || col == len(trees)-1 || row == len(trees[col])-1 {
		return true
	}

	tree := trees[row][col]

	blocked := false
	for i := 0; i < row; i++ {
		if trees[i][col].height >= tree.height {
			blocked = true
			break
		}
	}
	if !blocked {
		return true
	}

	blocked = false
	for i := len(trees) - 1; i > row; i-- {
		if trees[i][col].height >= tree.height {
			blocked = true
			break
		}
	}
	if !blocked {
		return true
	}

	blocked = false
	for i := len(trees[row]) - 1; i > col; i-- {
		if trees[row][i].height >= tree.height {
			blocked = true
			break
		}
	}
	if !blocked {
		return true
	}

	blocked = false
	for i := 0; i < col; i++ {
		if trees[row][i].height >= tree.height {
			blocked = true
			break
		}
	}
	if !blocked {
		return true
	}

	return false
}

func updateVisibleTrees(trees [][]*tree) int {
	numberVisibleTrees := 0
	for row := 0; row < len(trees); row++ {
		for col := 0; col < len(trees[row]); col++ {
			isVisible := isTreeVisible(trees, col, row)
			if isVisible {
				numberVisibleTrees++
			}
			trees[row][col].isVisible = isVisible
		}
	}

	return numberVisibleTrees
}

func getScenicScore(trees [][]*tree, row int, col int) int {
	if col == 0 || row == 0 || col == len(trees)-1 || row == len(trees[col])-1 {
		return 0
	}

	tree := trees[row][col]

	southScore := 0
	for i := row + 1; i < len(trees); i++ {
		southScore++
		if trees[i][col].height >= tree.height {
			break
		}
	}

	northScore := 0
	for i := row - 1; i >= 0; i-- {
		northScore++
		if trees[i][col].height >= tree.height {
			break
		}
	}

	eastScore := 0
	for i := col + 1; i < len(trees[row]); i++ {
		eastScore++
		if trees[row][i].height >= tree.height {
			break
		}
	}

	westScore := 0
	for i := col - 1; i >= 0; i-- {
		westScore++
		if trees[row][i].height >= tree.height {
			break
		}
	}

	return northScore * southScore * westScore * eastScore
}

func read() [][]*tree {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	trees := [][]*tree{}

	for scanner.Scan() {
		text := scanner.Text()
		treeLine := []*tree{}
		for _, c := range strings.Split(text, "") {
			i, _ := strconv.Atoi(c)
			treeLine = append(treeLine, &tree{height: i, isVisible: false})
		}
		trees = append(trees, treeLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return trees
}
