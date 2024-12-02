package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const MIN_DIFFERENCE = 1
const MAX_DIFFERENCE = 3

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func isLineSafe(numbers []int, correctedProblems int) bool {
	if correctedProblems > 1 {
		return false
	}
	if isLineSafe(numbers[1:], correctedProblems+1) {
		return true
	}
	if isLineSafe(numbers[0:len(numbers)-1], correctedProblems+1) {
		return true
	}
	rising := numbers[0] <= numbers[1]
	for ix := 1; ix < len(numbers); ix++ {
		difference := absDiff(numbers[ix-1], numbers[ix])
		if difference < MIN_DIFFERENCE || difference > MAX_DIFFERENCE || (rising && numbers[ix] < numbers[ix-1]) || (!rising && numbers[ix] > numbers[ix-1]) {
			if correctedProblems == 1 {
				return false
			}
			var slice []int
			if ix == 1 {
				slice = numbers[ix:]
			} else {
				slice = append(slice, numbers[0:ix-1]...)
				slice = append(slice, numbers[ix:]...)
			}
			if !isLineSafe(slice, correctedProblems+1) {
				slice = make([]int, 0)
				if ix+1 < len(numbers) {
					slice = append(slice, numbers[0:ix]...)
					slice = append(slice, numbers[ix+1:]...)
				} else {
					slice = append(slice, numbers[0:len(numbers)-1]...)
				}
				return isLineSafe(slice, correctedProblems+1)
			} else {
				return true
			}
		}
	}
	return true
}

func solve(file *os.File) uint32 {
	scanner := bufio.NewScanner(file)
	var numberOfSafe uint32 = 0

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		var parsedNumbers []int
		for _, number := range numbers {
			parsedNumber, _ := strconv.Atoi(number)
			parsedNumbers = append(parsedNumbers, parsedNumber)
		}
		if isLineSafe(parsedNumbers, 0) {
			numberOfSafe++
		} else {
			fmt.Println(line)
		}
	}

	return numberOfSafe
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(solve(file))
}
