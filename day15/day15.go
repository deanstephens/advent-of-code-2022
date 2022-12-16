package day15

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type sensor struct {
	pos                     utils.Coordinate
	closestBeacon           utils.Coordinate
	distanceToClosestBeacon int
}

func getDisplayAtPosition(coordinate utils.Coordinate, sensors []sensor, sensorFreeZoneMap map[string]bool) string {
	for _, s := range sensors {
		if s.pos.X == coordinate.X && s.pos.Y == coordinate.Y {
			return "S"
		}
		if s.closestBeacon.X == coordinate.X && s.closestBeacon.Y == coordinate.Y {
			return "B"
		}
	}
	_, ok := sensorFreeZoneMap[utils.Coordinate{X: coordinate.X, Y: coordinate.Y}.String()]
	if ok {
		return "#"
	} else {
		return "."
	}
}

func displayGrid(sensors []sensor, beaconFreeZoneMap map[string]bool) {
	fmt.Print("\033[H\033[2J")
	for y := -2; y <= 22; y++ {
		fmt.Printf("%2d ", y)
		for x := -2; x <= 25; x++ {
			fmt.Print(getDisplayAtPosition(utils.Coordinate{X: x, Y: y}, sensors, beaconFreeZoneMap))
		}
		fmt.Print("\n")
	}
}

func Part1() int {
	sensors := read(false)
	beaconFreeZoneMap := map[int]bool{}
	knownBeacons := map[string]bool{}

	for _, s := range sensors {
		scan := getScanZoneAtY(s, 2000000)
		for _, scanHit := range scan {
			beaconFreeZoneMap[scanHit] = true
		}
		knownBeacons[s.closestBeacon.String()] = true
	}
	count := len(beaconFreeZoneMap)

	for b, _ := range knownBeacons {
		if pos := strings.Split(b, ","); pos[1] == "2000000" {
			count--
		}
	}

	return count
}

func Part2() int {
	sensors := read(false)

	hit := utils.Coordinate{}

	for y := 0; y <= 4000000; y++ {
		for x := 0; x <= 4000000; x++ {
			sensorHit := false
			for _, s := range sensors {
				distance := s.distanceToClosestBeacon

				magX, _ := utils.GetMagnitudeAndDirection2d(s.pos.X, x)
				magY, _ := utils.GetMagnitudeAndDirection2d(s.pos.Y, y)
				distanceToPoint := magX + magY

				if distanceToPoint <= distance {
					sensorHit = true
					x += distance - distanceToPoint
					continue
				}
			}
			if !sensorHit {
				hit = utils.Coordinate{
					X: x,
					Y: y,
				}
				fmt.Println("Hit", x, y)
			}
		}
		//if y%100 == 0 {
		//	fmt.Println(y)
		//}
	}

	return hit.X*4000000 + hit.Y
}

func getScanZoneAtY(s sensor, y int) []int {
	magX, _ := utils.GetMagnitudeAndDirection2d(s.pos.X, s.closestBeacon.X)
	magY, _ := utils.GetMagnitudeAndDirection2d(s.pos.Y, s.closestBeacon.Y)
	distance := magX + magY

	overlappingX := []int{}

	if s.pos.Y < y && s.pos.Y+distance > y {
		for x := -distance + (y - s.pos.Y); x <= distance-(y-s.pos.Y); x++ {
			overlappingX = append(overlappingX, x+s.pos.X)
		}
		return overlappingX
	}
	if s.pos.Y > y && s.pos.Y-distance < y {
		for x := -distance - (y - s.pos.Y); x <= distance+(y-s.pos.Y); x++ {
			overlappingX = append(overlappingX, x+s.pos.X)
		}
		return overlappingX
	}

	return nil
}

func updateBeaconFreeZones(s sensor, beaconFreeZoneMap map[string]bool) {
	magX, _ := utils.GetMagnitudeAndDirection2d(s.pos.X, s.closestBeacon.X)
	magY, _ := utils.GetMagnitudeAndDirection2d(s.pos.Y, s.closestBeacon.Y)

	distance := magX + magY

	for y := 0; y <= distance; y++ {
		for x := -distance + y; x <= distance-y; x++ {
			beaconFreeZoneMap[utils.Coordinate{X: s.pos.X + x, Y: s.pos.Y + y}.String()] = true
		}
	}

	for y := 0; y >= -distance; y-- {
		for x := -distance - y; x <= distance+y; x++ {
			beaconFreeZoneMap[utils.Coordinate{X: s.pos.X + x, Y: s.pos.Y + y}.String()] = true
		}
	}
}

func read(testInput bool) []sensor {
	fileName := "input.txt"
	if testInput {
		fileName = "test.txt"
	}
	file, err := os.Open("day15/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sensors := []sensor{}

	for scanner.Scan() {
		text := scanner.Text()

		r := regexp.MustCompile(`Sensor at x=(?P<sensorX>-?\d+), y=(?P<sensorY>-?\d+): closest beacon is at x=(?P<beaconX>-?\d+), y=(?P<beaconY>-?\d+)`)
		matches := r.FindStringSubmatch(text)

		sensorX, _ := strconv.Atoi(matches[1])
		sensorY, _ := strconv.Atoi(matches[2])
		beaconX, _ := strconv.Atoi(matches[3])
		beaconY, _ := strconv.Atoi(matches[4])

		magX, _ := utils.GetMagnitudeAndDirection2d(sensorX, beaconX)
		magY, _ := utils.GetMagnitudeAndDirection2d(sensorY, beaconY)
		distance := magX + magY

		sensors = append(sensors, sensor{
			pos:                     utils.Coordinate{X: sensorX, Y: sensorY},
			closestBeacon:           utils.Coordinate{X: beaconX, Y: beaconY},
			distanceToClosestBeacon: distance,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sensors
}
