package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func mergeTwoNumbers(left, right uint64) uint64 {
	var cnt uint64 = 0
	rightCopy := right
	for rightCopy != 0 {
		cnt++
		rightCopy /= 10
	}
	power := uint64(math.Pow(10, float64(cnt)))
	return left*power + right
}

func isTargetReachable(current uint64, target uint64, numbers []uint64, operation rune) bool {
	if len(numbers) == 0 && current != target {
		return false
	}
	switch operation {
	case '+':
		current += numbers[0]
	case '*':
		current *= numbers[0]
	case '|':
		current = mergeTwoNumbers(current, numbers[0])
	}
	if current == target && len(numbers) == 1 {
		return true
	}
	if current > target && len(numbers) > 1 {
		return false
	}

	slice := numbers[1:]
	return isTargetReachable(current, target, slice, '+') || isTargetReachable(current, target, slice, '*') || isTargetReachable(current, target, slice, '|')
}

func solve(file *os.File) uint64 {
	scanner := bufio.NewScanner(file)

	var sum uint64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ": ")
		target := lineParts[0]
		numbers := strings.Split(lineParts[1], " ")
		intTarget, _ := strconv.ParseUint(target, 10, 0)
		var intNumbers []uint64
		for _, number := range numbers {
			intNumber, _ := strconv.ParseUint(number, 10, 0)
			intNumbers = append(intNumbers, intNumber)
		}
		if isTargetReachable(0, intTarget, intNumbers, '+') || isTargetReachable(0, intTarget, intNumbers, '*') {
			sum += intTarget
		}
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
