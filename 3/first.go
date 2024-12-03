package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func solve(line string) uint64 {
	mulRegex := regexp.MustCompile("mul\\([1-9][0-9]{0,2},[1-9][0-9]{0,2}\\)")
	digitsRegex := regexp.MustCompile("\\d+")

	matches := mulRegex.FindAllString(line, -1)
	var sum uint64 = 0

	for _, match := range matches {
		digits := digitsRegex.FindAllString(match, -1)
		a, _ := strconv.Atoi(digits[0])
		b, _ := strconv.Atoi(digits[1])

		sum += uint64(a * b)
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
