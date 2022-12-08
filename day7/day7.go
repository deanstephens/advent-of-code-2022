package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type directory struct {
	path            string
	parentDirectory *directory
	files           map[string]int
	directories     map[string]*directory
	totalSize       int
}

const MAX_FILE_SIZE_TO_FIND = 100000
const AVAILABLE_DISK = 70000000 - 30000000

func Part1() int {
	rootDir := read()
	recursivelyAddFileSizes(&rootDir)

	totalSmall := getTotalSmallDirectorySizes(&rootDir, MAX_FILE_SIZE_TO_FIND)

	return totalSmall
}

func Part2() int {
	rootDir := read()
	recursivelyAddFileSizes(&rootDir)

	requiredSpace := rootDir.totalSize - AVAILABLE_DISK
	closestSubDir := getClosestSubDirSizeOverMin(&rootDir, requiredSpace)

	return closestSubDir
}

func getTotalSmallDirectorySizes(dir *directory, size int) int {
	totalSmallSize := 0
	for _, nextDir := range dir.directories {
		totalSmallSize += getTotalSmallDirectorySizes(nextDir, size)
	}

	if dir.totalSize < size {
		totalSmallSize += dir.totalSize
	}

	return totalSmallSize
}

func getClosestSubDirSizeOverMin(dir *directory, min int) int {
	closestSizeForAll := -1
	for _, nextDir := range dir.directories {
		smallSize := getClosestSubDirSizeOverMin(nextDir, min)
		if smallSize != -1 && (smallSize < closestSizeForAll || closestSizeForAll == -1) {
			closestSizeForAll = smallSize
		}
	}

	if closestSizeForAll == -1 && dir.totalSize > min {
		closestSizeForAll = dir.totalSize
	}

	return closestSizeForAll
}

func recursivelyAddFileSizes(dir *directory) int {
	totalSize := 0
	for _, nextDir := range dir.directories {
		totalSize += recursivelyAddFileSizes(nextDir)
	}
	for _, file := range dir.files {
		totalSize += file
	}
	dir.totalSize = totalSize

	return totalSize
}

func read() directory {
	var currentDirectory *directory

	var rootDirectory = directory{
		path:            "/",
		parentDirectory: nil,
		files:           nil,
		directories:     nil,
		totalSize:       0,
	}

	currentDirectory = &rootDirectory

	file, err := os.Open("day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// optionally, resize scanner's capacity for lines over 64K, see next example https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	for scanner.Scan() {
		text := scanner.Text()
		commands := strings.Split(text, " ")

		switch commands[0] {
		case "$":
			if commands[1] == "ls" {

			}

			if commands[1] == "cd" {
				if commands[2] == ".." {
					currentDirectory = currentDirectory.parentDirectory
					break
				}
				if currentDirectory.directories != nil {
					nextDirectory, _ := currentDirectory.directories[commands[2]]
					currentDirectory = nextDirectory
				}
			}
			break
		case "dir":
			if currentDirectory.directories == nil {
				currentDirectory.directories = map[string]*directory{}
			}

			newDirectory := directory{
				path:            fmt.Sprintf("%s%s/", currentDirectory.path, commands[1]),
				parentDirectory: currentDirectory,
				files:           nil,
				directories:     nil,
				totalSize:       0,
			}
			currentDirectory.directories[commands[1]] = &newDirectory
			break
		default:
			if currentDirectory.files == nil {
				currentDirectory.files = map[string]int{}
			}
			fileSize, _ := strconv.Atoi(commands[0])
			currentDirectory.files[commands[1]] = fileSize
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rootDirectory
}
