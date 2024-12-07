package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const OBSTACLE = '#'

func read(file *os.File) [][]rune {
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input [][]rune = make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lineRunes := make([]rune, len(line))
		for ix, character := range line {
			lineRunes[ix] = character
		}
		input = append(input, lineRunes)
	}

	return input
}

func determineStartingPositionAndDirection(input [][]rune, guardSymbols map[rune][]int) (int, int, rune) {
	for ix, line := range input {
		for jx, character := range line {
			if _, exists := guardSymbols[character]; exists {
				return ix, jx, input[ix][jx]
			}
		}
	}

	return -1, -1, '$'
}

func step(input [][]rune, startX, startY int, guardDirection rune, guardSymbols map[rune][]int) (int, int, rune) {
	nextGuardDirections := map[rune]rune{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}
	stepValue := guardSymbols[guardDirection]
	nextX, nextY := startX+stepValue[0], startY+stepValue[1]
	nextGuardDirection := guardDirection

	if nextX >= len(input) || nextY >= len(input[0]) || nextX < 0 || nextY < 0 {
		return nextX, nextY, guardDirection
	}
	for input[nextX][nextY] == '#' {
		nextGuardDirection = nextGuardDirections[nextGuardDirection]
		stepValue := guardSymbols[nextGuardDirection]
		nextX, nextY = startX+stepValue[0], startY+stepValue[1]
	}
	return nextX, nextY, nextGuardDirection
}

func countVisited(input [][]rune) uint32 {
	var count uint32 = 0
	for _, line := range input {
		for _, character := range line {
			if character == 'X' {
				count++
			}
		}
	}

	return count
}

func solve(input [][]rune) uint32 {
	guardSymbols := map[rune][]int{
		'^': {-1, 0},
		'>': {0, 1},
		'v': {1, 0},
		'<': {0, -1},
	}
	xSize, ySize := len(input), len(input[0])
	currentX, currentY, guardDirection := determineStartingPositionAndDirection(input, guardSymbols)

	for currentX < xSize && currentY < ySize && currentX >= 0 && currentY >= 0 {
		input[currentX][currentY] = 'X'
		currentX, currentY, guardDirection = step(input, currentX, currentY, guardDirection, guardSymbols)
	}

	return countVisited(input)
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	input := read(file)

	fmt.Println(solve(input))
}
