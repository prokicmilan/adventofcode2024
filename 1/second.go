package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(file *os.File) uint64 {
	scanner := bufio.NewScanner(file)
	var leftColumn []int
	rightColumn := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		for ix, number := range numbers {
			intNumber, _ := strconv.Atoi(number)
			if ix == 0 {
				leftColumn = append(leftColumn, intNumber)
			} else {
				if reps, present := rightColumn[intNumber]; present {
					rightColumn[intNumber] = reps + 1
				} else {
					rightColumn[intNumber] = 1
				}
			}
		}
	}
	sum := uint64(0)
	for _, number := range leftColumn {
		sum += uint64(number * rightColumn[number])
	}

	return sum
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(solve(file))
}
