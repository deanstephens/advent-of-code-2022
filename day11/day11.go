package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type test struct {
	divisor     int64
	monkeyTrue  int
	monkeyFalse int
}

type monkey struct {
	items       []int64
	operation   func(old int64) int64
	test        test
	inspections int
}

type monkeyList []monkey

func (s monkeyList) Len() int {
	return len(s)
}
func (s monkeyList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s monkeyList) Less(i, j int) bool {
	return s[i].inspections < s[j].inspections
}

type destressFunction func(worry int64) int64

var itemString = ""
var itemStrings = map[string][]int64{}

func monkeyTurn(monkeys monkeyList, monkeyIndex int, destressFn destressFunction) {
	currentMonkey := monkeys[monkeyIndex]
	for _, item := range currentMonkey.items {
		worry := currentMonkey.operation(item)
		worry = destressFn(worry)
		itemString += strconv.Itoa(monkeyIndex)
		if worry%currentMonkey.test.divisor == 0 {
			monkeys[currentMonkey.test.monkeyTrue].items = append(monkeys[currentMonkey.test.monkeyTrue].items, worry)
		} else {
			monkeys[currentMonkey.test.monkeyFalse].items = append(monkeys[currentMonkey.test.monkeyFalse].items, worry)
		}
	}

	monkeys[monkeyIndex].inspections += len(currentMonkey.items)
	monkeys[monkeyIndex].items = []int64{}
}

func Part1() int {
	monkeys := read(false)

	destressFn := func(worry int64) int64 {
		return worry / 3
	}

	for round := 0; round < 20; round++ {
		for i, _ := range monkeys {
			monkeyTurn(monkeys, i, destressFn)
		}
		fmt.Println(monkeys)
	}

	sort.Sort(monkeys)

	return monkeys[len(monkeys)-1].inspections * monkeys[len(monkeys)-2].inspections
}

func Part2() int {
	monkeys := read(false)

	maxModulus := int64(1)
	for _, m := range monkeys {
		maxModulus *= m.test.divisor
	}

	destressFn := func(worry int64) int64 {
		return worry % maxModulus
	}

	for round := 1; round <= 10000; round++ {
		for i, _ := range monkeys {
			monkeyTurn(monkeys, i, destressFn)
		}
		if round == 1 || round == 20 || round%1000 == 0 {
			fmt.Println("== After round", round, "==")
			for i, m := range monkeys {
				fmt.Println("Monkey", i, "inspected items", m.inspections, "times.")
			}
			fmt.Println("")
		}
	}

	sort.Sort(monkeys)

	return monkeys[len(monkeys)-1].inspections * monkeys[len(monkeys)-2].inspections
}

type operand func(a int64, b int64) int64

var times operand = func(a int64, b int64) int64 {
	return a * b
}
var division operand = func(a int64, b int64) int64 {
	return a / b
}
var addition operand = func(a int64, b int64) int64 {
	return a + b
}
var subtraction operand = func(a int64, b int64) int64 {
	return a - b
}

func read(testInput bool) monkeyList {
	fileName := "input.txt"
	if testInput {
		fileName = "test.txt"
	}
	file, err := os.Open("day11/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	monkeys := monkeyList{}

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "Monkey") {
			monkeys = append(monkeys, monkey{})
		}

		if strings.HasPrefix(text, "  Starting items: ") {
			itemsString := strings.Replace(text, "  Starting items: ", "", 1)
			itemsSplit := strings.Split(itemsString, ", ")

			items := []int64{}
			for _, item := range itemsSplit {
				i, _ := strconv.Atoi(item)
				items = append(items, int64(i))
			}

			monkeys[len(monkeys)-1].items = items
		}

		if strings.HasPrefix(text, "  Operation: new = ") {
			operationString := strings.Replace(text, "  Operation: new = ", "", 1)
			operations := strings.Split(operationString, " ")
			monkeys[len(monkeys)-1].operation = func(old int64) int64 {
				var firstNum int64
				if operations[0] == "old" {
					firstNum = old
				} else {
					n, _ := strconv.Atoi(operations[0])
					firstNum = int64(n)
				}

				var secondNum int64
				if operations[2] == "old" {
					secondNum = old
				} else {
					n, _ := strconv.Atoi(operations[2])
					secondNum = int64(n)
				}

				var op operand
				switch operations[1] {
				case "*":
					op = times
					break
				case "/":
					op = division
					break
				case "+":
					op = addition
					break
				case "-":
					op = subtraction
					break
				}

				return op(firstNum, secondNum)
			}
		}

		if strings.HasPrefix(text, "  Test: divisible by ") {
			testString := strings.Replace(text, "  Test: divisible by ", "", 1)
			testDivisor, _ := strconv.Atoi(testString)

			monkeys[len(monkeys)-1].test = test{divisor: int64(testDivisor)}
		}

		if strings.HasPrefix(text, "    If true: throw to monkey ") {
			trueMonkeyString := strings.Replace(text, "    If true: throw to monkey ", "", 1)
			trueMonkey, _ := strconv.Atoi(trueMonkeyString)

			monkeys[len(monkeys)-1].test.monkeyTrue = trueMonkey
		}

		if strings.HasPrefix(text, "    If false: throw to monkey ") {
			falseMonkeyString := strings.Replace(text, "    If false: throw to monkey ", "", 1)
			falseMonkey, _ := strconv.Atoi(falseMonkeyString)

			monkeys[len(monkeys)-1].test.monkeyFalse = falseMonkey
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return monkeys
}
