package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(file *os.File) uint32 {
	scanner := bufio.NewScanner(file)

	mappingRestrictions := true
	restrictions := make([][]int, 100)
	var sum uint32 = 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			mappingRestrictions = false
			continue
		}
		if mappingRestrictions {
			numbers := strings.Split(line, "|")
			left, _ := strconv.Atoi(numbers[0])
			right, _ := strconv.Atoi(numbers[1])

			restrictions[right] = append(restrictions[right], left)

		} else {
			numbers := strings.Split(line, ",")
			bannedFollowers := make(map[int]bool)
			validUpdate := true
			for _, number := range numbers {
				intNumber, _ := strconv.Atoi(number)
				if isBanned := bannedFollowers[intNumber]; isBanned {
					validUpdate = false
					break
				}
				for _, bannedFollower := range restrictions[intNumber] {
					bannedFollowers[bannedFollower] = true
				}
			}
			if validUpdate {
				intNumber, _ := strconv.Atoi(numbers[(len(numbers)-1)/2])
				sum += uint32(intNumber)
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
