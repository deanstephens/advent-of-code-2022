package day13

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

type pair struct {
	first  interface{}
	second interface{}
}

type packets = []interface{}

func Part1() int {
	pairs := read(false)

	sumOfIndices := 0
	for i, p := range pairs {
		fmt.Printf("\n== Pair %d ==\n", i+1)
		comparison := compare(p, 0, true)
		if comparison == -1 {
			sumOfIndices += i + 1
		}
	}

	return sumOfIndices
}

func Part2() int {
	pairs := read(false)

	packetsList := []interface{}{}

	var dividerPacket1 interface{}
	json.Unmarshal([]byte("[[2]]"), &dividerPacket1)
	var dividerPacket2 interface{}
	json.Unmarshal([]byte("[[6]]"), &dividerPacket2)

	packetsList = append(packetsList, dividerPacket1, dividerPacket2)
	for _, p := range pairs {
		packetsList = append(packetsList, p.first, p.second)
	}

	sort.SliceStable(packetsList, func(i, j int) bool {
		p := pair{first: packetsList[i], second: packetsList[j]}
		return compare(p, 0, false) == -1
	})

	return findDecoderKey(dividerPacket1, dividerPacket2, packetsList)
}

func findDecoderKey(dividerPacket1 interface{}, dividerPacket2 interface{}, packetsList []interface{}) int {
	decoderKey := 1
	for i, p := range packetsList {
		isDivider1 := compare(pair{first: p, second: dividerPacket1}, 0, false) == 0
		isDivider2 := compare(pair{first: p, second: dividerPacket2}, 0, false) == 0
		if isDivider1 || isDivider2 {
			decoderKey *= (i + 1)
		}
	}
	return decoderKey
}

func printPadded(s string, depth int, logging bool) {
	if !logging {
		return
	}
	paddingString := fmt.Sprintf("%*s", depth*2, "")
	fmt.Printf("%s%s", paddingString, s)
}

func tryGetIntFromInterface(i interface{}) (int, bool) {
	asFloat, ok := i.(float64)
	if ok {
		return int(asFloat), true
	}
	asInt, ok := i.(int)
	return asInt, ok
}

func compare(p pair, depth int, logging bool) int {
	printPadded(fmt.Sprintf("- Compare %v vs %v\n", p.first, p.second), depth, logging)
	firstInt, fIsInt := tryGetIntFromInterface(p.first)
	secondInt, sIsInt := tryGetIntFromInterface(p.second)

	if fIsInt && sIsInt {
		if firstInt < secondInt {
			return -1
		} else if firstInt > secondInt {
			return 1
		} else {
			return 0
		}
	}

	firstArray, faok := p.first.([]interface{})
	if !faok {
		firstArray = []interface{}{}
		firstArray = append(firstArray, firstInt)
		if fIsInt {
			printPadded(fmt.Sprintf("- Mixed types; convert left to %v and retry comparison\n", firstArray), depth+1, logging)
			return compare(pair{first: firstArray, second: p.second.([]interface{})}, depth+1, logging)
		}
	}

	secondArray, saok := p.second.([]interface{})
	if !saok {
		secondArray = []interface{}{}
		secondArray = append(secondArray, secondInt)
		if sIsInt {
			printPadded(fmt.Sprintf("- Mixed types; convert right to %v and retry comparison\n", secondArray), depth+1, logging)
			return compare(pair{first: firstArray, second: secondArray}, depth+1, logging)
		}
	}

	for i := 0; i < len(firstArray); i++ {
		if len(secondArray) <= i {
			printPadded(fmt.Sprintf("- Right side ran out of items, so inputs are not in the right order\n"), depth+1, logging)
			return 1
		}
		comparison := compare(pair{first: firstArray[i], second: secondArray[i]}, depth+1, logging)
		if comparison != 0 {
			if comparison == -1 {
				printPadded(fmt.Sprintf("- Left side is smaller, so inputs are in the right order\n"), depth+2, logging)
			} else {
				printPadded(fmt.Sprintf("- Right side is smaller, so inputs are in the wrong order\n"), depth+2, logging)
			}
			return comparison
		}
	}

	if len(secondArray) > len(firstArray) {
		printPadded(fmt.Sprintf("- Left side ran out of items, so inputs are in the right order\n"), depth+1, logging)
		return -1
	}

	return 0
}

func read(testInput bool) []pair {
	fileName := "input.txt"
	if testInput {
		fileName = "test.txt"
	}
	file, err := os.Open("day13/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pairs := []pair{}

	for {
		newPair := pair{}

		canScan := scanner.Scan()
		if !canScan {
			break
		}
		pair1Text := scanner.Text()
		err := json.Unmarshal([]byte(pair1Text), &newPair.first)
		if err != nil {
			panic(err)
		}

		scanner.Scan()
		pair2Text := scanner.Text()
		err = json.Unmarshal([]byte(pair2Text), &newPair.second)
		if err != nil {
			panic(err)
		}

		scanner.Scan()

		pairs = append(pairs, newPair)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return pairs
}
