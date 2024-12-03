package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solve(line string) uint64 {
	mulRegex := regexp.MustCompile("mul\\([1-9][0-9]{0,2},[1-9][0-9]{0,2}\\)|do\\(\\)|don't\\(\\)")
	digitsRegex := regexp.MustCompile("\\d+")

	matches := mulRegex.FindAllString(line, -1)

	var sum uint64 = 0
	mulEnabled := true

	for _, match := range matches {
		if strings.Contains(match, "mul") && mulEnabled {
			digits := digitsRegex.FindAllString(match, -1)
			a, _ := strconv.Atoi(digits[0])
			b, _ := strconv.Atoi(digits[1])
			sum += uint64(a * b)
		}
		if strings.Contains(match, "don't") {
			mulEnabled = false
			continue
		}
		if strings.Contains(match, "do") {
			mulEnabled = true
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
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	line := scanner.Text()

	fmt.Println(solve(line))
}
