package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type MemoryBlock struct {
	fileId, size int
}

func solve(line string) uint64 {
	readingFile := true
	var fileId int = 0
	disk := make([]MemoryBlock, 0)
	for _, character := range line {
		number, _ := strconv.Atoi(string(character))

		memoryBlock := MemoryBlock{
			fileId: fileId,
			size:   number,
		}
		fileId++
		if !readingFile {
			memoryBlock.fileId = -1
			fileId--
		}
		readingFile = !readingFile
		disk = append(disk, memoryBlock)
	}
	left, right := 0, len(disk)-1
	upperLimit := len(disk) - 1
	for {
		for left < upperLimit && disk[left].fileId != -1 {
			left++
		}
		for right >= 0 && disk[right].fileId == -1 {
			right--
		}
		if right <= left || left >= len(disk) || right < 0 {
			break
		}
		disk[left].fileId = disk[right].fileId
		if disk[left].size > disk[right].size {
			emptyBlock := MemoryBlock{
				fileId: -1,
				size:   disk[left].size - disk[right].size,
			}
			disk = append(disk, emptyBlock)
			upperLimit++
			for ix := upperLimit - 1; ix > left; ix-- {
				disk[ix+1] = disk[ix]
				if ix == right {
					right++
				}
			}
			disk[left].size = disk[right].size
			disk[left+1] = emptyBlock
		}
		disk[right].size -= disk[left].size
		if disk[right].size == 0 {
			disk[right].fileId = -1
		}
	}

	var checksum uint64 = 0
	var index uint64 = 0
	for _, memoryBlock := range disk {
		if memoryBlock.fileId == -1 {
			break
		}
		for i := 0; i < memoryBlock.size; i++ {
			checksum += index * uint64(memoryBlock.fileId)
			index++
		}
	}

	return checksum
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	fmt.Println(solve(line))
}
