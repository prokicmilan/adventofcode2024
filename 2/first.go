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

func solve(file *os.File) uint32 {
	scanner := bufio.NewScanner(file)
	numberOfSafe := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		previousNumber := -1
		var rising, determinedMonotony, unsafe bool
		for _, number := range numbers {
			parsedNumber, _ := strconv.Atoi(number)
			if previousNumber == -1 {
				previousNumber = parsedNumber
				continue
			}
			if !determinedMonotony {
				determinedMonotony = true
				if parsedNumber > previousNumber {
					rising = true
				} else {
					rising = false
				}
			}
			if absDiff(parsedNumber, previousNumber) < MIN_DIFFERENCE || absDiff(parsedNumber, previousNumber) > MAX_DIFFERENCE || (rising && parsedNumber < previousNumber) || (!rising && parsedNumber > previousNumber) {
				unsafe = true
				break
			}
			previousNumber = parsedNumber
		}
		if !unsafe {
			numberOfSafe++
		}
	}

	return uint32(numberOfSafe)
}

func main2() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(solve(file))
}
