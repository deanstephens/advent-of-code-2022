package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type sack struct {
	firstHalf  []int
	secondHalf []int
	sharedItem int
}

func Day3Part1() int {
	sacks := read()
	fmt.Println(sacks)

	totalScore := 0

	for _, sack := range sacks {
		totalScore += sack.sharedItem
	}

	return totalScore
}

func Day3Part2() int {
	sacks := read()
	fmt.Println(sacks)

	totalScore := 0

	chunks := chunkSlice(sacks, 3)

	for c, chunk := range chunks {
		for i := 1; i <= 52; i++ {
			if doAllSacksHaveItem(i, chunk) {
				totalScore += i
				fmt.Println("Chunk", c, "has item", i)
			}
		}
	}

	return totalScore
}

func doAllSacksHaveItem(item int, sacks []sack) bool {
	for _, sack := range sacks {
		if !sack.doesSackHaveItem(item) {
			return false
		}
	}
	return true
}

func (s sack) doesSackHaveItem(item int) bool {
	for _, i := range s.firstHalf {
		if i == item {
			return true
		}
	}
	for _, i := range s.secondHalf {
		if i == item {
			return true
		}
	}
	return false
}

func chunkSlice(slice []sack, chunkSize int) [][]sack {
	var chunks [][]sack
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func read() []sack {
	sacks := []sack{}

	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	for scanner.Scan() {
		text := []rune(scanner.Text())

		firstHalf := []int{}
		secondHalf := []int{}
		shared := 0

		for i := 0; i < len(text); i++ {
			priority := 0
			if text[i] >= 65 && text[i] <= 90 {
				priority = int(text[i]) - 64 + 26
			} else {
				priority = int(text[i]) - 96
			}

			if i < len(text)/2 {
				firstHalf = append(firstHalf, priority)
			} else {
				for _, ele := range firstHalf {
					if ele == priority {
						fmt.Println("Shared", ele)
						shared = ele
					}
				}
				secondHalf = append(secondHalf, priority)
			}
		}
		sacks = append(sacks, sack{firstHalf: firstHalf, secondHalf: secondHalf, sharedItem: shared})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sacks
}
