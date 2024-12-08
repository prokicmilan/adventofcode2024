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

	for _, antennaPositions := range antennas {
		for ix := 0; ix < len(antennaPositions); ix++ {
			for jx := 0; jx < len(antennaPositions); jx++ {
				if ix == jx {
					continue
				}
				pointA := antennaPositions[ix]
				pointB := antennaPositions[jx]
				dx := pointB.x - pointA.x
				dy := pointB.y - pointA.y
				x := pointA.x
				y := pointA.y

				for 0 <= x && x < rows && 0 <= y && y < cols {
					antiNodes[Point{x, y}] = true
					x += dx
					y += dy
				}
			}
		}
	}

	// for antiNode := range antiNodes {
	// 	if antiNode.x >= 0 && antiNode.x < rows && antiNode.y >= 0 && antiNode.y < cols {
	// 		cnt++
	// 	}
	// }
	return uint32(len(antiNodes))
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(file))
}
