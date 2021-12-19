package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("day1!")
	depths := readInput()

	var sums []int

	sum := depths[0] + depths[1] + depths[2]

	sums = append(sums, sum)

	for i := 1; i < len(depths); i++ {

		if (i + 2) >= len(depths) {
			break
		}

		currSum := depths[i] + depths[i+1] + depths[i+2]

		sums = append(sums, currSum)

	}

	fmt.Println(findIncreasisity(sums))

}

func readInput() []int {
	var depths []int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		nonString, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, nonString)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return depths
}

func findIncreasisity(arr []int) int {

	count := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			count++
		}
	}

	return count
}
