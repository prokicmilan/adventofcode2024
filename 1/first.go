package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func absDiff(x, y uint64) uint64 {
	if x < y {
		return y - x
	}
	return x - y
}

func solve(file *os.File) uint64 {
	scanner := bufio.NewScanner(file)
	var leftColumn, rightColumn []int
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		for ix, stringNumber := range numbers {
			intNumber, _ := strconv.Atoi(stringNumber)
			if ix == 0 {
				leftColumn = append(leftColumn, intNumber)
			} else {
				rightColumn = append(rightColumn, intNumber)
			}
		}
	}
	sort.Slice(leftColumn, func(i, j int) bool {
		return leftColumn[i] < leftColumn[j]
	})
	sort.Slice(rightColumn, func(i, j int) bool {
		return rightColumn[i] < rightColumn[j]
	})
	sum := uint64(0)
	for ix := 0; ix < len(leftColumn); ix++ {
		diff := absDiff(uint64(leftColumn[ix]), uint64(rightColumn[ix]))
		sum += diff
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
