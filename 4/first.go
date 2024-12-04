package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func solve(file *os.File) uint32 {
	scanner := bufio.NewScanner(file)
	var input []string
	var sum uint32 = 0

	for scanner.Scan() {
		line := scanner.Text()

		input = append(input, line)
	}

	var word string
	for ix := 0; ix < len(input); ix++ {
		for jx := 0; jx < len(input[0]); jx++ {
			if rune(input[ix][jx]) == 'X' {
				// check right
				word = ""
				buffer := bytes.NewBufferString(word)
				for jxx := jx; jxx < len(input[0]) && jxx-jx <= 3; jxx++ {
					buffer.WriteByte(input[ix][jxx])
				}
				word = buffer.String()
				if strings.Compare(word, "XMAS") == 0 {
					sum++
				}
				// check left
				word = ""
				buffer = bytes.NewBufferString(word)
				for jxx := jx; jxx >= 0 && jx-jxx <= 3; jxx-- {
					buffer.WriteByte(input[ix][jxx])
				}
				word = buffer.String()
				if strings.Compare(word, "XMAS") == 0 {
					sum++
				}
				// check down
				word = ""
				buffer = bytes.NewBufferString(word)
				for ixx := ix; ixx < len(input) && ixx-ix <= 3; ixx++ {
					buffer.WriteByte(input[ixx][jx])
				}
				word = buffer.String()
				if strings.Compare(word, "XMAS") == 0 {
					sum++
				}
				// check up
				word = ""
				buffer = bytes.NewBufferString(word)
				for ixx := ix; ixx >= 0 && ix-ixx <= 3; ixx-- {
					buffer.WriteByte(input[ixx][jx])
				}
				word = buffer.String()
				if strings.Compare(word, "XMAS") == 0 {
					sum++
				}

				// check right diagonal down
				word = ""
				buffer = bytes.NewBufferString(word)
				for ixx, jxx := ix, jx; ixx < len(input) && jxx < len(input[0]) && ixx-ix <= 3; ixx, jxx = ixx+1, jxx+1 {
					buffer.WriteByte(input[ixx][jxx])
				}
				word = buffer.String()
				if strings.Compare(word, "XMAS") == 0 {
					sum++
				}

				// check right diagonal up
				word = ""
				buffer = bytes.NewBufferString(word)
				for ixx, jxx := ix, jx; ixx >= 0 && jxx < len(input) && ix-ixx <= 3; ixx, jxx = ixx-1, jxx+1 {
					buffer.WriteByte(input[ixx][jxx])
				}
				word = buffer.String()
				if strings.Compare(word, "XMAS") == 0 {
					sum++
				}

				// check left diagonal down
				word = ""
				buffer = bytes.NewBufferString(word)
				for ixx, jxx := ix, jx; ixx < len(input) && jxx >= 0 && ixx-ix <= 3; ixx, jxx = ixx+1, jxx-1 {
					buffer.WriteByte(input[ixx][jxx])
				}
				word = buffer.String()
				if strings.Compare(word, "XMAS") == 0 {
					sum++
				}

				// check left diagonal up
				word = ""
				buffer = bytes.NewBufferString(word)
				for ixx, jxx := ix, jx; ixx >= 0 && jxx >= 0 && ix-ixx <= 3; ixx, jxx = ixx-1, jxx-1 {
					buffer.WriteByte(input[ixx][jxx])
				}
				word = buffer.String()
				if strings.Compare(word, "XMAS") == 0 {
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
