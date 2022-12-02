package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	food     []int
	calories int
}

func Day1Part1() int {
	elfs := read()
	//fmt.Println(elfs)

	maxCalories := 0
	maxCaloriesElf := 0

	for i, elf := range elfs {
		fmt.Println(i, elf)
		if elf.calories > maxCalories {

			maxCalories = elf.calories
			maxCaloriesElf = i + 1
		}
	}

	return maxCaloriesElf
}

func Day1Part2() int {
	elfs := read()
	sort.Slice(elfs, func(i, j int) bool {
		return elfs[i].calories > elfs[j].calories
	})

	return elfs[0].calories + elfs[1].calories + elfs[2].calories
}

func read() []elf {
	elfs := []elf{{
		food:     []int{},
		calories: 0,
	}}

	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			elfs = append(elfs, elf{
				food:     []int{},
				calories: 0,
			})
			continue
		}

		num, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}

		elfs[len(elfs)-1].food = append(elfs[len(elfs)-1].food, num)
		elfs[len(elfs)-1].calories += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return elfs
}
