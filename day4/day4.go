package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cleaningRoster struct {
	min int
	max int
}

func Day4Part1() int {
	cleaningRosters := read()
	fmt.Println(cleaningRosters)

	totalScore := 0

	for _, cr := range cleaningRosters {
		o := getOverlap(cr[0], cr[1])

		overlapLength := o.max - o.min + 1
		aRangeLength := cr[0].max - cr[0].min + 1
		bRangeLength := cr[1].max - cr[1].min + 1

		fmt.Println(cr[0], cr[1], "overlap: ", overlapLength, "a : ", aRangeLength, "b : ", bRangeLength)
		if o.max == -1 {
			continue
		}

		if overlapLength == aRangeLength || overlapLength == bRangeLength {
			fmt.Println(overlapLength)
			totalScore += 1
		}
	}

	return totalScore
}

func Day4Part2() int {
	cleaningRosters := read()
	fmt.Println(cleaningRosters)

	totalScore := 0

	for _, cr := range cleaningRosters {
		o := getOverlap(cr[0], cr[1])

		overlapLength := o.max - o.min + 1
		aRangeLength := cr[0].max - cr[0].min + 1
		bRangeLength := cr[1].max - cr[1].min + 1

		fmt.Println(cr[0], cr[1], "overlap: ", overlapLength, "a : ", aRangeLength, "b : ", bRangeLength)
		if o.max == -1 {
			continue
		}

		if overlapLength > 0 {
			fmt.Println(overlapLength)
			totalScore += 1
		}
	}

	return totalScore
}

func getOverlap(a, b cleaningRoster) cleaningRoster {
	if a.min <= b.min {
		if a.max < b.min {
			return cleaningRoster{min: -1, max: -1}
		}
		if a.max <= b.max {
			return cleaningRoster{b.min, a.max}
		}
		return cleaningRoster{b.min, b.max}
	}

	if b.max < a.min {
		return cleaningRoster{min: -1, max: -1}
	}
	if b.max <= a.max {
		return cleaningRoster{a.min, b.max}
	}
	return cleaningRoster{a.min, a.max}
}

func read() [][]cleaningRoster {
	cleaningRosters := [][]cleaningRoster{}

	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	for scanner.Scan() {
		text := scanner.Text()

		s := strings.Split(text, ",")

		first := strings.Split(s[0], "-")
		f1, _ := strconv.Atoi(first[0])
		f2, _ := strconv.Atoi(first[1])

		second := strings.Split(s[1], "-")
		s1, _ := strconv.Atoi(second[0])
		s2, _ := strconv.Atoi(second[1])

		cleaningRosters = append(cleaningRosters, []cleaningRoster{
			{min: f1, max: f2},
			{min: s1, max: s2},
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return cleaningRosters
}
