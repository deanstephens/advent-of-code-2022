package day16

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	combin "gonum.org/v1/gonum/stat/combin"
)

type valve struct {
	name            string
	flowRate        int
	connectedValves []connection
}

type connection struct {
	connectedValve *valve
	cost           int
}

func Part1() int {
	valves := read(true)
	fmt.Println(valves)

	shortestPaths := map[string]int{}

	for _, v := range valves {
		if v.flowRate != 0 || v.name == "AA" {
			shortestPath(v, v, []*valve{}, shortestPaths)
		}
	}

	valveNodes := map[string]*valve{}
	totalFlowRate := 0
	for _, v := range valves {
		if v.flowRate != 0 || v.name == "AA" {
			v.connectedValves = []connection{}
			valveNodes[v.name] = v
			for _, vDest := range valves {
				if vDest.flowRate != 0 && vDest.name != v.name {
					connectionCost := shortestPaths[v.name+"_"+vDest.name]
					v.connectedValves = append(v.connectedValves, connection{connectedValve: vDest, cost: connectionCost})
				}
			}
			totalFlowRate += v.flowRate
		}
	}

	bestPath := 0
	navigate(valveNodes["AA"], 31, []*valve{}, []state{}, &bestPath)

	return 0
}

func Part2() int {
	valves := read(false)
	fmt.Println(valves)

	shortestPaths := map[string]int{}

	for _, v := range valves {
		if v.flowRate != 0 || v.name == "AA" {
			shortestPath(v, v, []*valve{}, shortestPaths)
		}
	}

	rootNode := &valve{}
	valveNodes := []*valve{}
	totalFlowRate := 0
	for _, v := range valves {
		if v.name == "AA" {
			rootNode = v
		}
		if v.flowRate != 0 {
			valveNodes = append(valveNodes, v)
		}
		if v.flowRate != 0 || v.name == "AA" {
			v.connectedValves = []connection{}
			for _, vDest := range valves {
				if vDest.flowRate != 0 && vDest.name != v.name {
					connectionCost := shortestPaths[v.name+"_"+vDest.name]
					v.connectedValves = append(v.connectedValves, connection{connectedValve: vDest, cost: connectionCost})
				}
			}
			totalFlowRate += v.flowRate
		}
	}

	bestCombo := 0
	for i := 1; i < len(valveNodes)-1; i++ {
		combinations := combin.Combinations(len(valveNodes)-1, i)
		for _, c := range combinations {
			preOpenedValves1 := []*valve{}
			preOpenedValves2 := []*valve{}
			for j, v := range valveNodes {
				if utils.IndexOf(c, j) != -1 {
					preOpenedValves1 = append(preOpenedValves1, v)
				} else {
					preOpenedValves2 = append(preOpenedValves2, v)
				}
			}
			bestPath1 := 0
			bestPath2 := 0
			p1 := navigate(rootNode, 27, preOpenedValves1, []state{}, &bestPath1)
			p2 := navigate(rootNode, 27, preOpenedValves2, []state{}, &bestPath2)

			if p1+p2 > bestCombo {
				bestCombo = p1 + p2
			}
			fmt.Println(p1, p2)
		}
	}

	return bestCombo
}

func shortestPath(start *valve, current *valve, traversedPath []*valve, shortestPaths map[string]int) int {
	if sp, ok := shortestPaths[start.name+"_"+current.name]; ok && sp < len(traversedPath) {
		return -1
	}

	if start.name != current.name {
		shortestPaths[start.name+"_"+current.name] = len(traversedPath)
	}

	distance := -1
	for _, cv := range current.connectedValves {
		sp := shortestPath(start, cv.connectedValve, append(traversedPath, cv.connectedValve), shortestPaths)

		if sp < distance || distance == -1 {
			distance = sp
		}
	}

	return distance
}

type state struct {
	minute           int
	valveName        string
	flowChange       int
	pressureChange   int
	pressureReleased int
	potentialLost    int
}

func navigate(v *valve, minutesRemaining int, openValves []*valve, states []state, bestPath *int) (pressureChange int) {
	if minutesRemaining <= 0 || utils.IndexOf(openValves, v) != -1 {
		totalPressure := 0
		checkedStates := states[:len(states)-1]
		for _, s := range checkedStates {
			if s.pressureChange > 0 {
				totalPressure += s.pressureChange
			}
		}
		if totalPressure > *bestPath {
			*bestPath = totalPressure
			return totalPressure
		}
		return -1
	}
	currentOpenedValves := openValves
	if v.flowRate != 0 {
		currentOpenedValves = append(openValves, v)
	}

	bestPressureChange := -1

	for _, cv := range v.connectedValves {
		s := state{}
		s.flowChange = cv.connectedValve.flowRate
		s.minute = 30 - (minutesRemaining - (cv.cost + 2))
		s.valveName = cv.connectedValve.name
		s.pressureChange = (minutesRemaining - (cv.cost + 2)) * cv.connectedValve.flowRate

		pc := navigate(cv.connectedValve, minutesRemaining-(cv.cost+1), currentOpenedValves, append(states, s), bestPath)
		if pc > bestPressureChange {
			bestPressureChange = pc
		}
	}

	return bestPressureChange
}

func read(testInput bool) map[string]*valve {
	fileName := "input.txt"
	if testInput {
		fileName = "test.txt"
	}
	file, err := os.Open("day16/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	valves := map[string]*valve{}
	valveConnections := map[string][]string{}

	for scanner.Scan() {
		text := scanner.Text()

		r := regexp.MustCompile(`Valve (?P<valveName>\w+) has flow rate=(?P<flowRate>\d+); tunnel(s){0,1} lead(s){0,1} to valve(s){0,1} (?P<connectedValves>\w+(, \w+)*)`)
		matches := utils.FindNamedMatches(r, text)

		flowRate, _ := strconv.Atoi(matches["flowRate"])
		valves[matches["valveName"]] = &valve{name: matches["valveName"], flowRate: flowRate}
		valveConnections[matches["valveName"]] = strings.Split(matches["connectedValves"], ", ")
	}

	for v, connections := range valveConnections {
		for _, c := range connections {
			valves[v].connectedValves = append(valves[v].connectedValves, connection{connectedValve: valves[c], cost: 1})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return valves
}
