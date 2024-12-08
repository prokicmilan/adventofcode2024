package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x, y int
}

func determineAntennaPositions(file *os.File) (map[rune][]Point, int, int) {
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var antennas = make(map[rune][]Point)
	row := 0
	cols := 0

	for scanner.Scan() {
		line := scanner.Text()
		cols = len(line)
		for jx, character := range line {
			if character != '.' {
				point := Point{row, jx}
				if _, exists := antennas[character]; !exists {
					antennas[character] = make([]Point, 0)
				}
				arr := antennas[character]
				arr = append(arr, point)
				antennas[character] = arr
			}
		}
		row++
	}

	return antennas, row, cols
}

func solve(file *os.File) uint32 {
	antennas, rows, cols := determineAntennaPositions(file)
	var antiNodes = make(map[Point]bool)
	var cnt uint32 = 0

	for _, antennaPositions := range antennas {
		for ix := 0; ix < len(antennaPositions); ix++ {
			for jx := ix + 1; jx < len(antennaPositions); jx++ {
				pointA := antennaPositions[ix]
				pointB := antennaPositions[jx]
				antiNodePosition := Point{
					2*pointA.x - pointB.x,
					2*pointA.y - pointB.y,
				}
				antiNodes[antiNodePosition] = true
				antiNodePosition = Point{
					2*pointB.x - pointA.x,
					2*pointB.y - pointA.y,
				}
				antiNodes[antiNodePosition] = true
			}
		}
	}

	for antiNode := range antiNodes {
		if antiNode.x >= 0 && antiNode.x < rows && antiNode.y >= 0 && antiNode.y < cols {
			cnt++
		}
	}
	return cnt
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(file))
}
