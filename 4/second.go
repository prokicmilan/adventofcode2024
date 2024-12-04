package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func solve(file *os.File) uint32 {
	scanner := bufio.NewScanner(file)
	var input []string
	var sum uint32 = 0

	for scanner.Scan() {
		line := scanner.Text()

		input = append(input, line)
	}

	for ix := 1; ix < len(input)-1; ix++ {
		for jx := 1; jx < len(input[0])-1; jx++ {
			if rune(input[ix][jx]) == 'A' {
				if ((rune(input[ix-1][jx-1]) == 'M' && rune(input[ix+1][jx+1]) == 'S') || (rune(input[ix-1][jx-1]) == 'S' && rune(input[ix+1][jx+1]) == 'M')) && ((rune(input[ix-1][jx+1]) == 'M' && rune(input[ix+1][jx-1]) == 'S') || (rune(input[ix-1][jx+1]) == 'S' && rune(input[ix+1][jx-1]) == 'M')) {
					sum++
				}
			}
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
