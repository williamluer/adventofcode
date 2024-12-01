package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func allZeros(numbers []int) bool {
	for _, num := range numbers {
		if num != 0 {
			return false
		}
	}
	return true
}

func findNext(previousSequences []int, currentSequences []int) int {
	if allZeros(currentSequences) {
		return previousSequences[0]
	}
	var newSequence []int
	for i := 0; i < len(currentSequences)-1; i++ {
		newSequence = append(newSequence, currentSequences[i+1]-currentSequences[i])
	}
	if len(previousSequences) > 0 {
		return previousSequences[0] - findNext(currentSequences, newSequence)
	}
	return findNext(currentSequences, newSequence)
}

func main() {
	filename := "9.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reg := regexp.MustCompile(`-?\d+`)
	scanner := bufio.NewScanner(file)

	var histories [][]int
	for scanner.Scan() {
		current_line := scanner.Text()
		seqStrs := reg.FindAllString(current_line, -1)
		var seqInts []int
		for _, o := range seqStrs {
			v, _ := strconv.Atoi(o)
			seqInts = append(seqInts, v)
		}
		histories = append(histories, seqInts)
	}

	var placeholder []int
	sumVals := 0
	for _, sequence := range histories {
		sumVals += findNext(placeholder, sequence)
	}
	fmt.Println(sumVals)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
